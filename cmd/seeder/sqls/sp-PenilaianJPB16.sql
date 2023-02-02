DROP FUNCTION IF EXISTS "public"."penilaian_jpb16"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar);
CREATE OR REPLACE FUNCTION "public"."penilaian_jpb16"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
    kls_jpb16         char(1);
    nilai_dbkb_jpb16  numeric := 0;
    nilai_jpb16       numeric := 0;
    nilai_jpb         numeric := 0;
    kd_jpb            varchar(10);

    luas_bng          numeric := 0;
    kondisi_bng       char(1);
    jml_lantai_bng    int := 0;
    komp_utama        numeric := 0;
    nilai_fas1        numeric := 0;
    nilai_fas2        numeric := 0;
    nilai_fas3        numeric := 0;
    nilai1            numeric := 0;
    nilai_sebelum_susut numeric := 0;
    nilai_setelah_susut numeric := 0;
    thn_dibangun_bng  char(4);
    thn_renovasi_bng  char(4);
    nilai_susut        numeric := 0;
    penyusutan         numeric := 0;
begin
  -- cari data bangunan
  BEGIN
    SELECT "Jpb_Kode", "TahunDibangun", "TahunRenovasi", "LuasBangunan", "Kondisi", "JmlLantaiBng"
    INTO kd_jpb, thn_dibangun_bng, thn_renovasi_bng, luas_bng, kondisi_bng, jml_lantai_bng
    FROM "ObjekPajakBangunan"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" = NoUrut AND
      "JenisOp" = JenisOp AND
      "NoBangunan" = NoBangunan AND
      "Jpb_Kode" = '14';
  END;

  IF found THEN
    -- cari data kelas
    BEGIN
      SELECT "KelasBangunan16"
      INTO kls_jpb16
      FROM "Jpb16"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Kecamatan_Kode" = Kecamatan_Kode AND
          "Kelurahan_Kode" = Kelurahan_Kode AND
          "Blok_Kode" = Blok_Kode AND
          "NoUrut" = NoUrut AND
          "JenisOp" = JenisOp AND
          "NoBangunan" = NoBangunan;

    END;

    -- cari nilai dbkb
    BEGIN
      SELECT "NilaiDbkbJpb16"
      INTO nilai_dbkb_jpb16
      FROM "DbkbJpb16"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "TahunDbkbJpb16" = Tahun AND
          "KelasDbkbJpb16" = kls_jpb16 AND
          "LantaiMinJpb16" <= jml_lantai_bng AND
          "LantaiMaxJpb16" >= jml_lantai_bng;

    END;
    komp_utama := nilai_dbkb_jpb16;
    
    BEGIN
      SELECT "fasilitas_susut_x_luas"
      INTO nilai_fas1
      FROM Fasilitas_Susut_X_Luas(
        Provinsi_Kode,
        Daerah_Kode,
        Kecamatan_Kode,
        Kelurahan_Kode,
        Blok_Kode,
        NoUrut,
        JenisOp,
        NoBangunan,
        kd_jpb,
        kls_jpb16,
        Tahun
      );

    END;
    
    nilai1 := (komp_utama + nilai_fas1) * luas_bng;

    -- cari nilai fasilitas yang disusutkan
    BEGIN
      SELECT "fasilitas_susut"
      INTO nilai_fas2
      FROM Fasilitas_Susut(
        Provinsi_Kode,
        Daerah_Kode,
        Kecamatan_Kode,
        Kelurahan_Kode,
        Blok_Kode,
        NoUrut,
        JenisOp,
        NoBangunan,
        jml_lantai_bng,
        Tahun
      );

    END;
    
    nilai_sebelum_susut := nilai1 + nilai_fas2;

    -- cari prosentase penyusutan
    BEGIN
      SELECT "susut"
      INTO nilai_susut
      FROM Susut(
        CAST(Tahun AS int),
        CAST(thn_dibangun_bng AS int),
        CAST(thn_renovasi_bng AS int),
        kondisi_bng,
        nilai_sbl_susut,
        luas_bng,
        0
      );

    END;
    
    IF (nilai_susut IS NOT NULL) OR (nilai_susut = 0) THEN
        penyusutan    := ROUND(nilai_susut/100, 2);
        nilai_setelah_susut := (nilai_sbl_susut - (nilai_sbl_susut * penyusutan));
    ELSE
        nilai_setelah_susut := nilai_sbl_susut;
    END IF; 

    -- cari nilai fasilitas yang tidak disusutkan
    BEGIN
      SELECT "fasilitas_tidak_susut"
      INTO nilai_fas3
      FROM Fasilitas_Tidak_Susut(
        Provinsi_Kode,
        Daerah_Kode,
        Kecamatan_Kode,
        Kelurahan_Kode,
        Blok_Kode,
        NoUrut,
        JenisOp,
        NoBangunan,
        Tahun
      );

    END;
    
    nilai_jpb16 := nilai_stl_susut + nilai_fas3;
  ELSE 
    nilai_jpb16 := 0;
  END IF;    

  nilai_jpb := nilai_jpb16;

  return nilai_jpb;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
DROP FUNCTION IF EXISTS "public"."penilaian_jpb4"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar);
CREATE OR REPLACE FUNCTION "public"."penilaian_jpb4"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_jpb numeric := 0;
  kls_jpb4 char(1);
  nilai_dbkb_jpb4	numeric := 0;
  kd_jpb char(2);
  thn_dibangun char(4);
  thn_renovasi char(4);
  luas_bng numeric := 0;
  kondisi_bng char(1);
  jml_lantai_bng int := 1;
  komp_utama 		numeric := 0;
  nilai_fas1 		numeric := 0;
  nilai_fas2 		numeric := 0;
  nilai_fas3 		numeric := 0;
  nilai1     		numeric := 0;
  nilai_sbl_susut  numeric := 0;
  nilai_stl_susut  numeric := 0;
  besar_susut      numeric := 0;
  persen_susut     numeric := 0;

begin
  
  SELECT "Jpb_Kode", "LuasBangunan", "TahunDibangun", "TahunRenovasi", "Kondisi", "JmlLantaiBng"
  INTO kd_jpb, luas_bng, thn_dibangun, thn_renovasi, kondisi_bng, jml_lantai_bng
  FROM   "ObjekPajakBangunan"
  WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" = NoUrut AND
      "JenisOp" = JenisOp	AND
      "NoBangunan" = NoBangunan;

  if found then
    BEGIN
      SELECT "KelasBangunan4"
      INTO kls_jpb4
      FROM   "Jpb4"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
        "Daerah_Kode" = Daerah_Kode AND
        "Kecamatan_Kode" = Kecamatan_Kode AND
        "Kelurahan_Kode" = Kelurahan_Kode AND
        "Blok_Kode" = Blok_Kode AND
        "NoUrut" = NoUrut AND
        "JenisOp" = JenisOp AND
        "NoBangunan" = no_bng;
        
    END;

    BEGIN
      SELECT "NilaiDbkbJpb4"
      INTO nilai_dbkb_jpb4
      FROM "DbkbJpb4"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
        "Daerah_Kode" = Daerah_Kode AND
        "TahunDbkbJpb4" = Tahun AND
        "KelasDbkbJpb4" = kls_jpb4 AND
        "LantaiMinJpb4" <= jml_lantai_bng AND
        "LantaiMaxJpb4" >= jml_lantai_bng;

      if not found then
        nilai_dbkb_jpb4 = 0;
      end if;
      
    END;
    komp_utama := nilai_dbkb_jpb4;
    
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
        kls_jpb4,
        Tahun
      );

    END;

    nilai1 := (komp_utama + nilai_fas1) * luas_bng;
    
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

    nilai_sbl_susut := nilai1 + nilai_fas2;
    
    BEGIN
      SELECT "susut"
      INTO besar_susut
      FROM Susut(
        CAST(Tahun AS int),
        CAST(tahun_dibangun AS int),
        CAST(tahun_renovasi AS int),
        kondisi_bng,
        nilai_sbl_susut,
        luas_bng,
        0
      );

    END;

    IF (besar_susut IS NOT NULL) OR (besar_susut = 0) THEN
      persen_susut    := round(besar_susut/100, 2);
      nilai_stl_susut := (nilai_sbl_susut - (nilai_sbl_susut * persen_susut));
    ELSE
      nilai_stl_susut := nilai_sbl_susut;
    END IF;

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

    nilai_jpb := nilai_stl_susut + nilai_fas3;
  ELSE
    nilai_jpb := 0;
  end if;


  return nilai_jpb;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
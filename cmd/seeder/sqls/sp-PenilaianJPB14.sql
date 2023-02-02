DROP FUNCTION IF EXISTS "public"."penilaian_jpb14"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar);
CREATE OR REPLACE FUNCTION "public"."penilaian_jpb14"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  penyusutan numeric := 0;
  thn_dibangun_bng char(4);
  thn_renovasi_bng char(4);
  luas_bng numeric := 0;
  jml_lantai_bng int := 1;
  kondisi_bng char(1);
  nilai_jpb14 numeric := 0;
  nilai_jpb numeric := 0;
  nilai_dbkb_jpb14 numeric := 0;
  nilai_fasilitas  numeric := 0;
  nilai_susut      numeric := 0;
  nilai_tidak_susut numeric := 0;

begin
  -- cari data bangunan
  BEGIN
    SELECT a."TahunDibangun", a."TahunRenovasi", a."LuasBangunan", a."Kondisi", a."JmlLantaiBng"
    INTO thn_dibangun_bng, thn_renovasi_bng, luas_bng, kondisi_bng, jml_lantai_bng
    FROM "ObjekPajakBangunan" AS a
    WHERE a."Provinsi_Kode" = Provinsi_Kode AND
      a."Daerah_Kode" = Daerah_Kode AND
      a."Kecamatan_Kode" = Kecamatan_Kode AND
      a."Kelurahan_Kode" = Kelurahan_Kode AND
      a."Blok_Kode" = Blok_Kode AND
      a."NoUrut" = NoUrut AND
      a."JenisOp" = JenisOp AND
      a."NoBangunan" = NoBangunan AND
      a."Jpb_Kode" = '14';
  END;

  -- cari nilai komponen utama / m2
  BEGIN
    SELECT "NilaiDbkbJpb14"
    INTO nilai_dbkb_jpb14
    FROM "DbkbJpb14"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "TahunDbkbJpb14" = Tahun;
  END;

  nilai_jpb14 := nilai_dbkb_jpb14 * luas_bng;

  -- cari nilai fasilitas yang disusutkan
  BEGIN
    SELECT "fasilitas_susut"
    INTO nilai_fasilitas
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
  
  nilai_jpb14     := nilai_jpb14 + nilai_fasilitas;

  -- cari prosentase penyusutan
  BEGIN
    SELECT "susut"
    INTO nilai_susut
    FROM Susut(
      CAST(Tahun AS int),
      CAST(thn_dibangun_bng AS int),
      CAST(thn_renovasi_bng AS int),
      kondisi_bng,
      nilai_jpb14,
      luas_bng,
      0
    );

  END;
  
  -- cari nilai setelah disusutkan
  penyusutan := ROUND((nilai_susut * nilai_jpb14) / 100);
  nilai_jpb14 := nilai_jpb14 - penyusutan;

  -- cari nilai fasilitas yang tidak disusutkan
  BEGIN
    SELECT "fasilitas_tidak_susut"
    INTO nilai_tidak_susut
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

  nilai_jpb14 := nilai_jpb14 + nilai_tidak_susut;
  nilai_jpb := nilai_jpb14;

  return nilai_jpb;
  
end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
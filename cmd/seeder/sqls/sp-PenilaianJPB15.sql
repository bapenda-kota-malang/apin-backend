DROP FUNCTION IF EXISTS "public"."penilaian_jpb15"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar);
CREATE OR REPLACE FUNCTION "public"."penilaian_jpb15"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  luas_bng numeric := 0;
  kondisi_bng char(1);
  jml_lantai_bng int := 1;
  thn_dibangun_bng char(4);
  thn_renovasi_bng char(4);
  nilai_susut       numeric := 0;
  penyusutan        numeric := 0;

  letak_tangki_jpb15   char(1);
  jns_tangki_jpb15   char(1);
  kapasitas_tangki   numeric := 0; -- bigint
  nilai_dbkb         numeric := 0;
  nilai_jpb15        numeric := 0;
  nilai_jpb          numeric := 0;

  var_tahun        integer; 
  tahun_dibangun 	 integer;
  tahun_renovasi 	 integer;
  umur_efektif		 integer := 0;

begin
  -- cari data bangunan
  BEGIN
    SELECT "TahunDibangun", "TahunRenovasi", "LuasBangunan", "Kondisi"
    INTO thn_dibangun_bng, thn_renovasi_bng, luas_bng, kondisi_bng
    FROM "ObjekPajakBangunan"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" = NoUrut AND
      "JenisOp" = JenisOp AND
      "NoBangunan" = NoBangunan;

  END;

  IF found THEN
    BEGIN
      SELECT "LetakTanki, KapasitasTanki"
      INTO letak_tangki_jpb15, kapasitas_tangki
      FROM "Jpb15"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Kecamatan_Kode" = Kecamatan_Kode AND
          "Kelurahan_Kode" = Kelurahan_Kode AND
          "Blok_Kode" = Blok_Kode AND
          "NoUrut" = NoUrut AND
          "JenisOp" = JenisOp AND
          "NoBangunan" = NoBangunan;

    END;

    -- cari nilai tanki
    BEGIN
      SELECT "NilaiDbkbJpb15"
      INTO nilai_dbkb
      FROM "DbkbJpb15"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "TahunDbkbJpb15" = Tahun AND
          "JenisTangkiDbkbJpb15" = jns_tangki_jpb15 AND
          "KapasitasMinDbkbJpb15" <= kapasitas_tangki AND
          "KapasitasMaxDbkbJpb15" >= kapasitas_tangki;

    END;

    -- cari prosentase penyusutan
    var_tahun      := intval(Tahun);
    tahun_dibangun := intval(thn_dibangun_bng);
    tahun_renovasi := intval(thn_renovasi_bng);

    --- mencari umur efektif
    IF tahun_dibangun > 0 THEN
      -- jika tahun dibangun ada
      IF tahun_renovasi > 0 THEN
        -- jika tahun renovasi ada
        IF (var_tahun - tahun_renovasi) > 10 THEN
          -- (jika tahun pajak - tahun renovasi) > 10
          umur_efektif := ROUND(((var_tahun - tahun_dibangun) + (2*10)) / 3);
        ELSIF (var_tahun - tahun_renovasi) <= 10 THEN
          -- (jika tahun pajak - tahun renovasi) <= 10
          umur_efektif := ROUND(((var_tahun - tahun_dibangun) + (2*(var_tahun - tahun_renovasi))) / 3);
        ELSE
          umur_efektif := 0;
        END IF;

      ELSE
        -- tahun renovasi kosong
        IF (var_tahun - tahun_dibangun) > 10 THEN
          umur_efektif := ROUND(((tahun - tahun_dibangun) + (2*10)) / 3);
        ELSIF (var_tahun - tahun_dibangun) <= 10 THEN
          umur_efektif := var_tahun - tahun_dibangun;
        ELSE
          umur_efektif := 0;
        END IF;

      END IF;
    ELSE
      umur_efektif := 0;
    END IF;

    IF umur_efektif > 40 THEN
      umur_efektif := 40;
    END IF;

    nilai_susut := umur_efektif * 5;

    -- cari nilai jpb 15
    IF (nilai_susut IS NOT NULL) OR (nilai_susut = 0) THEN
        IF nilai_susut > 50 THEN
            nilai_susut := 50;
        END IF;
            penyusutan  := ROUND((nilai_susut / 100),2);
            nilai_jpb15 := (nilai_dbkb - (nilai_dbkb * penyusutan));
    ELSE
        nilai_jpb15 	 := nilai_dbkb;
    END IF;
  ELSE 
    nilai_jpb15 := 0;
  END IF;    

  nilai_jpb := nilai_jpb15;

  return nilai_jpb;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
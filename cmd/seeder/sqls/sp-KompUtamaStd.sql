DROP FUNCTION IF EXISTS "public"."komp_utama_std"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar, "jpb_kode" bpchar, "luasbangunan" int8, "jml_lantai_bng" int4, "jeniskonstruksi" bpchar);
CREATE OR REPLACE FUNCTION "public"."komp_utama_std"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar, "jpb_kode" bpchar, "luasbangunan" int8, "jml_lantai_bng" int4, "jeniskonstruksi" bpchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  komp_utama numeric := 0;
  kayu_ulin int;
 	nilai_dbkb numeric := 0;
 	Jpb char(2);
  tipe_bng char(5);
  kd_bng_lantai char(8);
  rec_tipe_bng record;
  c_tipe_bng cursor for SELECT "Tipe_Bng", "Bng_Lantai_Kode"
    FROM "BangunanLantai", "TipeBangunan"
    WHERE "Jpb_Kode" = jpb AND
      "LantaiMinBng" <= jml_lantai_bng AND
      "LantaiMaxBng" >= jml_lantai_bng AND
      "LuasMinTipe" <= LuasBangunan AND
      "LuasMaxTipe" >= LuasBangunan AND
      "Bng_Tipe" = "Tipe_Bng";

begin

  BEGIN
    SELECT "StatusKayuUlin"
    INTO kayu_ulin
    FROM "KayuUlin"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "ThnStatusKayuUlin" = Tahun;

    if not found then
      kayu_ulin := 0;
    end if;
  END;

  IF Jpb_Kode IN ('04','09') THEN
    Jpb := '02';
  ELSE
    Jpb := Jpb_Kode;
  END IF;

  -- proses c_fas1
  open c_tipe_bng;

  loop
    fetch c_tipe_bng into rec_tipe_bng;
    -- exit when no more row to fetch
    exit when not found;

    if found then
      BEGIN
        SELECT "NilaiDbkbStandard"
        INTO nilai_dbkb
        FROM "DbkbStandard"
        WHERE "Provinsi_Kode" = Provinsi_Kode
          AND "Daerah_Kode" = Daerah_Kode
          AND "ThnDbkbStandard" = Tahun
          AND "Jpb_Kode" = Jpb
          AND "TipeBng" = tipe_bng
          AND "Bng_Lantai_Kode" = rec_tipe_bng.kd_bng_lantai;

      if not found then
        nilai_dbkb := 0;
      end if;
      END;

      IF (JenisKonstruksi = '4') AND (kayu_ulin = '0') THEN
        komp_utama := (nilai_dbkb * 7/10);
      ELSE
        komp_utama := nilai_dbkb;
      END IF;
    end if;
  
  end loop;
  
  close c_tipe_bng;

  komp_utama := 0;

  return komp_utama;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
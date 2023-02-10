CREATE OR REPLACE FUNCTION "public"."komp_material"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  komp_material numeric := 0;
  jml_lantai_bng int := 1;
  jns_konstruksi char(1);
  jns_atap char(1);
  kd_dinding char(1);
  kd_lantai char(1);
  kd_langit char(1);
  atap numeric := 0;
  nilai_atap numeric := 0;
  nilai_dinding numeric := 0;
  nilai_lantai numeric := 0;
  nilai_langit numeric := 0;

begin

  SELECT "JenisKonstruksi", "JenisAtap", "KodeDinding", "KodeLantai", "KodeLangitLangit", "JmlLantaiBng"
	INTO jns_konstruksi,
    jns_atap,
    kd_dinding,
    kd_lantai,
    kd_langit,
    jml_lantai_bng
	FROM	"ObjekPajakBangunan"
	WHERE "Provinsi_Kode" = Provinsi_Kode AND
    "Daerah_Kode" = Daerah_Kode AND
    "Kecamatan_Kode" = Kecamatan_Kode AND
    "Kelurahan_Kode" = Kelurahan_Kode AND
    "Blok_Kode" = Blok_Kode AND
    "NoUrut" = NoUrut AND
    "JenisOp" = JenisOp	AND
    "NoBangunan" = NoBangunan;

  if found then
    -- cari nilai atap
    IF (jns_atap IS NOT NULL) OR (jns_atap = '0') THEN
      BEGIN
        SELECT "NilaiDbkbMaterial"
        INTO atap
        FROM "DbkbMaterial"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Thn_Dbkb_Material" = Tahun AND
          "Pekerjaan_Kode" = '23' AND
          substr("Kegiatan_Kode",2,1) = jns_atap;

        nilai_atap := (atap / jml_lantai_bng);

        if not found then
          nilai_atap := 0;
        end if;
      END;
    END IF;

    -- cari nilai dinding
    IF (kd_dinding IS NOT NULL) OR (kd_dinding = '0') THEN
		  IF kd_dinding = '1' THEN kd_dinding := '1';
		  ELSIF kd_dinding = '2' THEN kd_dinding := '2';
		  ELSIF kd_dinding = '3' THEN kd_dinding := '3';
		  ELSIF kd_dinding = '4' THEN kd_dinding := '7';
		  ELSIF kd_dinding = '5' THEN kd_dinding := '8';
		  ELSIF kd_dinding = '6' THEN kd_dinding := null;
		  END IF;

		  BEGIN
  			SELECT "NilaiDbkbMaterial"
  			INTO nilai_dinding
  			FROM "DbkbMaterial"
  			WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Thn_Dbkb_Material" = Tahun AND
          "Pekerjaan_Kode" = '23' AND
          substr("Kegiatan_Kode",2,1) = jns_atap;
  	    
        if not found then
          nilai_dinding := 0;
        end if;
      END;
    END IF;

    -- cari nilai lantai
    IF (kd_lantai IS NOT NULL) OR (kd_lantai = '0') THEN
      BEGIN
        SELECT "NilaiDbkbMaterial"
        INTO nilai_lantai
        FROM "DbkbMaterial"
  			WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Thn_Dbkb_Material" = Tahun AND
          "Pekerjaan_Kode" = '22' AND
          substr("Kegiatan_Kode",2,1) = kd_lantai;

        if not found then
          nilai_lantai := 0;
        end if; 
      END;
    END IF;

    -- cari nilai langit-langit
    IF (kd_langit IS NOT NULL) OR (kd_langit = '0') THEN
      BEGIN
        SELECT "NilaiDbkbMaterial"
        INTO nilai_langit
        FROM "DbkbMaterial"
  			WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Thn_Dbkb_Material" = Tahun AND
          "Pekerjaan_Kode" = '24' AND
          substr("Kegiatan_Kode",2,1) = kd_langit;
        
        if not found then
          nilai_langit := 0;
        end if;
      END;
    END IF;
    
    -- hitung nilai komponen material
    komp_material := (nilai_atap + nilai_dinding + nilai_lantai + nilai_langit);
  ELSE
    komp_material := 0;
  end if;

  return komp_material;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
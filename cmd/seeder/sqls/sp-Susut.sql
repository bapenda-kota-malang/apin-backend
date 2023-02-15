CREATE OR REPLACE FUNCTION "public"."susut"("tahun" int4, "tahun_dibangun" int4, "tahun_renovasi" int4, "kondisi_bangunan" bpchar, "nilai" numeric, "luas_bangunan" numeric, "flag_standard" numeric)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  persentase_susut numeric := 0;
  nilai_nir numeric := 0;
	umur_efektif numeric := 0;
	biaya_pengganti_baru numeric := 0;
	kd_range_penyusutan varchar(6);
begin

  -- mencari umur efektif
  IF Flag_Standard = 0 THEN
    -- jika bangunan non standart
    IF Tahun_Dibangun > 0 THEN
      -- jika tahun dibangun ada
      IF Tahun_Renovasi > 0 THEN
        -- jika tahun renovasi ada
        IF (Tahun - Tahun_Renovasi) > 10 THEN
          -- (jika tahun pajak - tahun renovasi) > 10
          umur_efektif := ROUND(((Tahun - Tahun_Dibangun) +	(2*10)) / 3);
        ELSE
          umur_efektif := ROUND(((Tahun - Tahun_Dibangun) + (2*(Tahun - Tahun_Renovasi))) / 3);
		    END IF;
      ELSE
        -- tahun renovasi kosong
        IF (Tahun - Tahun_Dibangun) > 10 THEN
          umur_efektif := ROUND(((Tahun - Tahun_Dibangun) + (2*10)) / 3);
        ELSE
          umur_efektif := Tahun - Tahun_Dibangun;
        END IF;
      END IF;
    ELSE
      RETURN 0;
    END IF;
  ELSE
    -- jika bangunan standart
    IF Tahun_Renovasi > 0 THEN
      umur_efektif := Tahun - Tahun_Renovasi;
    ELSE
      umur_efektif := Tahun - Tahun_Dibangun;
    END IF;
  END IF;

  IF umur_efektif > 40 THEN
    umur_efektif := 40;
	END IF;

  -- mencari biaya pengganti baru / m2
  IF Luas_Bangunan = 0 THEN
    Luas_Bangunan := 1;
	ELSE
    Luas_Bangunan := Luas_Bangunan;
	END IF;

	biaya_pengganti_baru := (Nilai * 1000) / Luas_Bangunan;

  -- mencari kode range penyusutan
  BEGIN
		SELECT "Range_Penyusutan_Kode"
		INTO kd_range_penyusutan
		FROM "RangePenyusutan"
		WHERE "NilaiMinPenyusutan" <  biaya_pengganti_baru AND "NilaiMaxPenyusutan" >= biaya_pengganti_baru;

    if not found then
      kd_range_penyusutan = 0;
    end if;
	END;

	-- mencari prosentase penyusutan
	BEGIN
    SELECT "NilaiPenyusutan"
    INTO persentase_susut
    FROM "Penyusutan"
    WHERE "UmurEfektif" = umur_efektif AND "Range_Penyusutan_Kode" = kd_range_penyusutan AND "KondisiBangunanSusut" = Kondisi_Bangunan;

    if not found then
      persentase_susut = 0;
    end if;
	END;

  return persentase_susut;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
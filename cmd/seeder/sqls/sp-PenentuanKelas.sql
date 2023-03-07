CREATE OR REPLACE FUNCTION "public"."penentuan_kelas"(IN "vln_flag_jenis" int8, IN "vlc_thn_pajak" varchar, IN "vln_nilai_perm2" int8, OUT "vlc_kd_kelas" bpchar, OUT "vlc_thn_awal_kls" bpchar, OUT "vln_njop_perm2" int8)
  RETURNS "pg_catalog"."record" AS $BODY$
declare
  vln_nilai_para numeric := 0;
begin

  --------------------------------------------------------------------------------
  --nilai per m2 di tabel dalam ribuan
  --------------------------------------------------------------------------------
  vln_nilai_para := vln_nilai_perm2/1000;

  --------------------------------------------------------------------------------
  -- jika flag jenis = 1 berarti untuk kelas tanah
  -- jika flag jenis = 2 berarti untuk kelas bangunan
  --------------------------------------------------------------------------------
  IF vln_nilai_perm2 > 0 THEN
	  IF vln_flag_jenis = 1 THEN
	    BEGIN
	      SELECT "KdTanah", "NilaiPerM2Tanah", "TahunAwalKelasTanah"
        INTO vlc_kd_kelas, vln_njop_perm2, vlc_thn_awal_kls
        FROM "KelasTanah"
	      WHERE  CAST("TahunAwalKelasTanah" AS int) <= CAST(vlc_thn_pajak AS int)
	        AND  CAST("TahunAkhirKelasTanah" AS int) >= CAST(vlc_thn_pajak AS int)
	        AND  "NilaiMinTanah"     <  vln_nilai_para
	        AND  "NilaiMaxTanah"     >= vln_nilai_para
          AND  "KdTanah" NOT IN ('XXX', '00');
	      EXCEPTION
	        WHEN OTHERS THEN
            vln_njop_perm2   := vln_nilai_para;
            vlc_kd_kelas     := 'XXX';
            vlc_thn_awal_kls := '1986';
	    END;
	  ELSIF vln_flag_jenis = 2 THEN
	    BEGIN
	      SELECT "KdBangunan", "NilaiPerM2Bangunan", "TahunAwalKelasBangunan"
		    INTO vlc_kd_kelas, vln_njop_perm2, vlc_thn_awal_kls
	      FROM "KelasBangunan"
	      WHERE CAST("TahunAwalKelasBangunan" AS int) <= CAST(vlc_thn_pajak AS int)
	        AND CAST("TahunAkhirKelasBangunan" AS int) >= CAST(vlc_thn_pajak AS int)
	        AND  "NilaiMinBangunan"     <  vln_nilai_para
	        AND  "NilaiMaxBangunan"     >= vln_nilai_para
			    AND  "KdBangunan" NOT IN ('XXX', '00');
	      EXCEPTION
          WHEN OTHERS THEN
            vln_njop_perm2   := vln_nilai_para;
            vlc_kd_kelas     := 'XXX';
            vlc_thn_awal_kls := '1986';
	    END;
	  END IF;

    --------------------------------------------------------------------------------
    --nilai per m2 di tabel dalam ribuan
    --------------------------------------------------------------------------------
	  vln_njop_perm2 := vln_njop_perm2 * 1000;
	  --vln_njop_perm2 := vln_njop_perm2;

  ELSE vln_njop_perm2   := 0;
	   vlc_kd_kelas     := 'XXX';
	   vlc_thn_awal_kls := '1986';
  END IF;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
CREATE OR REPLACE FUNCTION "public"."njoptkp_individu"("v_kd_propinsi" varchar, "v_kd_dati2" varchar, "v_kd_kecamatan" varchar, "v_kd_kelurahan" varchar, "v_kd_blok" varchar, "v_no_urut" varchar, "v_kd_jns_op" varchar, "v_thn_proses" bpchar)
  RETURNS "pg_catalog"."void" AS $BODY$
declare
  vlc_subjek_pajak_id       bigint;
	vln_kd_status_cabang      bigint;
	vln_total_luas_bng        bigint;
	vln_njop_bumi             bigint;
	vln_njop_bng              bigint;
	vlc_jns_bumi              char(1);
	vlc_kd_jpb                char(2);
	vlc_kd_propinsi           char(2);
	vlc_kd_dati2              char(2);
	vlc_kd_kecamatan          char(3);
	vlc_kd_kelurahan          char(3);
	vlc_kd_blok               char(3);
	vlc_no_urut               char(4);
	vlc_kd_jns_op             char(1);
	vlc_kd_propinsi_max       char(2);
	vlc_kd_dati2_max          char(2);
	vlc_kd_kecamatan_max      char(3);
	vlc_kd_kelurahan_max      char(3);
	vlc_kd_blok_max           char(3);
	vlc_no_urut_max           char(4);
	vlc_kd_jns_op_max         char(1);
	vlc_kd_propinsi_sp        char(2);
	vlc_kd_dati2_sp           char(2);
	vlc_kd_kecamatan_sp       char(3);
	vlc_kd_kelurahan_sp       char(3);
	vlc_kd_blok_sp            char(3);
	vlc_no_urut_sp            char(4);
	vlc_kd_jns_op_sp          char(1);
	vln_njop_bumi_beban       bigint;
	vln_njop_bng_beban        bigint;
	vld_tgl_cetak             date;
	vln_jml_bng               numeric := 0;
	vln_jml_op_bersama        numeric := 0;
	vln_max_non_op_bersama    numeric := 0;
	vln_max_op_bersama        numeric := 0;
	vln_proses                numeric := 0;  -- 0 = tidak, 1 = ya
	vln_njopmax               numeric := 0;
	vln_jml_op                numeric := 0;
	vln_jml_sp                numeric := 0;
	vln_tanahkosongall        numeric := 1;  -- 0 = tidak, 1 = ya
	vln_total_njop            numeric := 0;
	vln_ceksp                 numeric := 0;
	vlc_thn_proses            char(4) := v_thn_proses;

  rec_list_nop record;
  cursor_list_nop cursor for 
    SELECT "Provinsi_Kode", "Daerah_Kode", "Kecamatan_Kode", "Kelurahan_Kode", "Blok_Kode", "NoUrut", "JenisOp", "StatusCabang", "NjopBumi", "NjopBangunan"
    FROM "ObjekPajakPbb"
    WHERE "WajibPajakPbb_Id" = vlc_subjek_pajak_id;
begin

  ----------------------------------------------------------------
  --- Hapus dat_subjek_pajak_njotkp per NOP untuk menghindari double NJOPTKP (17 Juli 2003)
  -----------------------------------------------------------------
  BEGIN
    delete from "SubjekPajakNJOPTKP"
      where "Provinsi_Kode"  = v_kd_propinsi
        AND "Daerah_Kode"     = v_kd_dati2
        AND "Kecamatan_Kode"  = v_kd_kecamatan
        AND "Kelurahan_Kode"  = v_kd_kelurahan
        AND "Blok_Kode"       = v_kd_blok
        AND "NoUrut"          = v_no_urut
        AND "JenisOp"         = v_kd_jns_op;
      exception
        when others then null;
  END;

  -----------------------------------------------------------------
  --ambil subjek_pajak_id untuk nop yang bersangkutan
  -----------------------------------------------------------------
  SELECT "WajibPajakPbb_Id"
  INTO   vlc_subjek_pajak_id
  FROM   "ObjekPajakPbb"
  WHERE "Provinsi_Kode"  = v_kd_propinsi
    AND "Daerah_Kode"    = v_kd_dati2
    AND "Kecamatan_Kode" = v_kd_kecamatan
    AND "Kelurahan_Kode" = v_kd_kelurahan
    AND "Blok_Kode"      = v_kd_blok
    AND "NoUrut"         = v_no_urut
    AND "JenisOp"        = v_kd_jns_op;

  -----------------------------------------------------------------
  --cari jumlah op yang dimiliki subjek_pajak_id
  -----------------------------------------------------------------
  SELECT count(*)
  INTO   vln_jml_op
  FROM 	 "ObjekPajakPbb"
  WHERE  "WajibPajakPbb_Id" = vlc_subjek_pajak_id;

  -----------------------------------------------------------------
  -- jika SP hanya memiliki 1 OP
  -----------------------------------------------------------------
	IF vln_jml_op = 1 THEN
      -----------------------------------------------------------------
      -- ambil status cabang
      -- status cabang = 0 - bukan cabang, 1 - cabang
      -----------------------------------------------------------------
	  BEGIN
	    SELECT "Provinsi_Kode", "Daerah_Kode", "Kecamatan_Kode", "Kelurahan_Kode", "Blok_Kode", "NoUrut", "JenisOp", "StatusCabang"
      INTO vlc_kd_propinsi, vlc_kd_dati2, vlc_kd_kecamatan, vlc_kd_kelurahan, vlc_kd_blok, vlc_no_urut, vlc_kd_jns_op, vln_kd_status_cabang
	    FROM "ObjekPajakPbb"
      WHERE "WajibPajakPbb_Id" = vlc_subjek_pajak_id;
	    EXCEPTION
	      WHEN OTHERS THEN
		    vln_kd_status_cabang := 1;
    END;

    -----------------------------------------------------------------
    -- ambil jenis bumi
	  -- jenis bumi = 1 - tanah ada bangunan, 2 - kavling siap bangun,
	  --              3 - tanah kosong, 4 - fasilitas umum
    -----------------------------------------------------------------
	  BEGIN
      SELECT "JenisBumi"
	    INTO vlc_jns_bumi
      FROM "ObjekPajakBumi"
      WHERE "Provinsi_Kode"   = vlc_kd_propinsi
        AND "Daerah_Kode"		  = vlc_kd_dati2
        AND "Kecamatan_Kode"  = vlc_kd_kecamatan
        AND "Kelurahan_Kode"	= vlc_kd_kelurahan
        AND "Blok_Kode"		= vlc_kd_blok
        AND "NoUrut"		= vlc_no_urut
        AND "JenisOp"	= vlc_kd_jns_op
        AND "NoBumi"		= 1;
      EXCEPTION
	      WHEN OTHERS THEN
        vlc_jns_bumi := '4';
    END;

    -----------------------------------------------------------------
    -- jika tanah ada bangunan dan bukan status cabang
    -----------------------------------------------------------------
    IF (vlc_jns_bumi = '1') AND (vln_kd_status_cabang = 0) AND (vlc_kd_jns_op != '8') AND (vlc_kd_jns_op != '9') THEN
		  vln_proses := 1;
		  vlc_kd_propinsi_max     := vlc_kd_propinsi;
      vlc_kd_dati2_max        := vlc_kd_dati2;
      vlc_kd_kecamatan_max    := vlc_kd_kecamatan;
      vlc_kd_kelurahan_max    := vlc_kd_kelurahan;
      vlc_kd_blok_max         := vlc_kd_blok;
      vlc_no_urut_max         := vlc_no_urut;
      vlc_kd_jns_op_max       := vlc_kd_jns_op;
    ELSE
      vln_proses := 0;
    END IF;
  
  ELSE
    -----------------------------------------------------------------
	  -- jika SP memiliki lebih dari 1 OP
    -----------------------------------------------------------------
	  vln_njopmax          := 0;
	  vln_tanahkosongall   := 1;

    open cursor_list_nop;
    loop
      fetch cursor_list_nop into rec_list_nop;
      -- exit when no more row to fetch
      exit when not found;

      BEGIN
        vlc_kd_propinsi := rec_list_nop."Provinsi_Kode";
        vlc_kd_dati2 := rec_list_nop."Daerah_Kode";
        vlc_kd_kecamatan := rec_list_nop."Kecamatan_Kode";
        vlc_kd_kelurahan := rec_list_nop."Kelurahan_Kode";
        vlc_kd_blok := rec_list_nop."Blok_Kode";
        vlc_no_urut := rec_list_nop."NoUrut";
        vlc_kd_jns_op := rec_list_nop."JenisOp";
        vln_kd_status_cabang := rec_list_nop."StatusCabang";
        vln_njop_bumi := rec_list_nop."NjopBumi";
        vln_njop_bng := rec_list_nop."NjopBangunan";
      END;

      -----------------------------------------------------------------
      -- ambil jenis bumi
      -----------------------------------------------------------------
      BEGIN
        SELECT "JenisBumi"
        INTO vlc_jns_bumi
        FROM "ObjekPajakBumi"
        WHERE "Provinsi_Kode"   = vlc_kd_propinsi
          AND "Daerah_Kode"		  = vlc_kd_dati2
          AND "Kecamatan_Kode"  = vlc_kd_kecamatan
          AND "Kelurahan_Kode"	= vlc_kd_kelurahan
          AND "Blok_Kode"		= vlc_kd_blok
          AND "NoUrut"		= vlc_no_urut
          AND "JenisOp"	= vlc_kd_jns_op
          AND "NoBumi"		= 1;
        EXCEPTION
          WHEN OTHERS THEN
          vlc_jns_bumi := '4';
      END;

      IF (vlc_jns_bumi != '4') AND (vln_kd_status_cabang = 0) AND (vlc_kd_jns_op != '8') AND (vlc_kd_jns_op != '9') THEN
        IF vlc_jns_bumi = '1' THEN
          vln_tanahkosongall   := 0;
        END IF;
        
        vln_total_njop := vln_njop_bumi + vln_njop_bng;
        
        -----------------------------------------------------------------
        -- cek anggota op bersama
        -----------------------------------------------------------------
        BEGIN
          SELECT count(*)
          INTO vln_jml_op_bersama
          FROM "AnggotaObjekPajak"
          WHERE "Provinsi_Kode"  = vlc_kd_propinsi
            AND "Daerah_Kode"    = vlc_kd_dati2
            AND "Kecamatan_Kode" = vlc_kd_kecamatan
            AND "Kelurahan_Kode" = vlc_kd_kelurahan
            AND "Blok_Kode"      = vlc_kd_blok
            AND "NoUrut"         = vlc_no_urut
            AND "JenisOp"        = vlc_kd_jns_op;
          EXCEPTION
            WHEN OTHERS THEN
            vln_jml_op_bersama := 0;
        END;

        IF vln_jml_op_bersama > 0 THEN
          BEGIN
            SELECT njop_bumi_beban, njop_bng_beban
            INTO vln_njop_bumi_beban, vln_njop_bng_beban
            FROM  "AnggotaObjekPajak"
            WHERE "Provinsi_Kode"  = vlc_kd_propinsi
              AND "Daerah_Kode"    = vlc_kd_dati2
              AND "Kecamatan_Kode" = vlc_kd_kecamatan
              AND "Kelurahan_Kode" = vlc_kd_kelurahan
              AND "Blok_Kode"      = vlc_kd_blok
              AND "NoUrut"         = vlc_no_urut
              AND "JenisOp"        = vlc_kd_jns_op;
            EXCEPTION
              WHEN OTHERS THEN
              vln_njop_bumi_beban := 0;
              vln_njop_bng_beban  := 0;
          END;
          
          vln_total_njop := vln_total_njop + vln_njop_bumi_beban + vln_njop_bng_beban;
        END IF;

        -----------------------------------------------------------------
        -- ambil NOP yang total njop-nya paling besar
        -----------------------------------------------------------------
        IF (vln_total_njop >= vln_njopmax) THEN
          vlc_kd_propinsi_max     := vlc_kd_propinsi;
          vlc_kd_dati2_max        := vlc_kd_dati2;
          vlc_kd_kecamatan_max    := vlc_kd_kecamatan;
          vlc_kd_kelurahan_max    := vlc_kd_kelurahan;
          vlc_kd_blok_max         := vlc_kd_blok;
          vlc_no_urut_max         := vlc_no_urut;
          vlc_kd_jns_op_max       := vlc_kd_jns_op;
          vln_njopmax			 :=	vln_total_njop;
        END IF;
      END IF;
    
    end loop;
    close cursor_list_nop;

    IF vln_tanahkosongall = 0 THEN
      vln_proses := 1;
	  ELSE
      vln_proses := 0;
	  END IF;

  END IF;

  IF vln_proses = 1 THEN
    -----------------------------------------------------------------
    -- cek data di "SubjekPajakNJOPTKP"
    -----------------------------------------------------------------
    BEGIN
      SELECT count(*)
      INTO vln_ceksp
		  FROM "SubjekPajakNJOPTKP"
		  WHERE "Id" = vlc_subjek_pajak_id;
	    EXCEPTION
	      WHEN OTHERS THEN
        vln_ceksp := 0;
    END;

    IF vln_ceksp = 0 THEN
      INSERT INTO "SubjekPajakNJOPTKP"(
        "Id",
		    "Provinsi_Kode",
        "Daerah_Kode",
        "Kecamatan_Kode",
        "Kelurahan_Kode",
        "Blok_Kode",
        "NoUrut",
        "JenisOp",
        "TahunPajak")
      VALUES(vlc_subjek_pajak_id,
        vlc_kd_propinsi_max,
        vlc_kd_dati2_max,
        vlc_kd_kecamatan_max,
        vlc_kd_kelurahan_max,
        vlc_kd_blok_max,
        vlc_no_urut_max,
        vlc_kd_jns_op_max,
        v_thn_proses);
    ELSE
      -----------------------------------------------------------------
      -- cek apakah nop pada "SubjekPajakNJOPTKP"
		  -- sudah dicetak sppt-nya
      -----------------------------------------------------------------
      BEGIN
        SELECT "Provinsi_Kode", "Daerah_Kode", "Kecamatan_Kode", "Kelurahan_Kode", "Blok_Kode", "NoUrut", "JenisOp"
        INTO   vlc_kd_propinsi_sp, vlc_kd_dati2_sp, vlc_kd_kecamatan_sp, vlc_kd_kelurahan_sp, vlc_kd_blok_sp,
          vlc_no_urut_sp, vlc_kd_jns_op_sp
        FROM "SubjekPajakNJOPTKP"
        WHERE "Id" = vlc_subjek_pajak_id;
        EXCEPTION
          WHEN OTHERS THEN
		      null;
      END;

      BEGIN
        SELECT "TanggalCetak_sppt"
			  INTO vld_tgl_cetak
			  FROM "Sppt"
			  WHERE "Propinsi_Id"    = vlc_kd_propinsi_sp 
          AND "Dati2_Id"       = vlc_kd_dati2_sp
				  AND "Kecamatan_Id"   = vlc_kd_kecamatan_sp
				  AND "Keluarahan_Id"   = vlc_kd_kelurahan_sp
				  AND "Blok_Id"        = vlc_kd_blok_sp
				  AND "NoUrut"        = vlc_no_urut_sp
				  AND "JenisOP_Id"      = vlc_kd_jns_op_sp
				  AND "TahunPajakskp_sppt" = v_thn_proses;
		    EXCEPTION
		      WHEN OTHERS THEN
          vld_tgl_cetak := null;
		  END;

		  IF vld_tgl_cetak IS NOT NULL THEN
        UPDATE "SubjekPajakNJOPTKP"
          SET "Provinsi_Kode"  = vlc_kd_propinsi_max,
            "Daerah_Kode" 	 = vlc_kd_dati2_max,
			     	"Kecamatan_Kode" = vlc_kd_kecamatan_max,
            "Kelurahan_Kode" = vlc_kd_kelurahan_max,
            "Blok_Kode" 	 = vlc_kd_blok_max,
            "NoUrut" 	 = vlc_no_urut_max,
            "JenisOp" 	 = vlc_kd_jns_op_max,
            "TahunPajak"  = v_thn_proses
          WHERE "Id" = vlc_subjek_pajak_id;
      END IF;
    END IF;
	ELSE
    ------------------------------------------------------------------------------
	  -- seharusnya tidak diproses, bila ada maka hapus [Teguh - Rachmat 20/12/2000]
    ------------------------------------------------------------------------------
    BEGIN
      DELETE FROM "SubjekPajakNJOPTKP"
		  WHERE "Id" = vlc_subjek_pajak_id;
	    EXCEPTION
	   	  WHEN OTHERS THEN NULL;
    END;
	END IF;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
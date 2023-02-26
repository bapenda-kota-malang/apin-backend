create or replace function Penetapan_Tunggal_SPPT(
  plc_nop varchar(2),
  plc_tahun char(4),
	plc_kd_kanwil_bank char(2),
  plc_kd_kppbb_bank char(2),
	plc_bank_tunggal char(2),
	plc_bank_persepsi char(2),
	plc_bank_tp char(2),
  plc_tgl_jatuh_tempo date,
  plc_tgl_terbit date,
  plc_tgl_cetak date,
  plc_nip_pencetak_sppt varchar(10),
	pln_luas_bumi bigint,
	pln_luas_bng bigint,
	pln_luas_bumi_beban bigint,
	pln_luas_bng_beban bigint,
	pln_njop_bumi bigint,
	pln_njop_bng bigint,
	pln_njop_bumi_beban bigint,
	pln_njop_bng_beban bigint,
	pln_param_sekarang bigint
)
returns void
language plpgsql    
as $$
declare
	vln_njop_bng_btkp numeric := 0;
	vln_temp_luas_bumi numeric := 0;
	vln_temp_luas_bng numeric := 0;
	vln_temp_njop_bumi numeric := 0;
	vln_temp_njop_bng numeric := 0;
	vlc_subjek_pajak_id numeric := 0;
	vlc_no_persil numeric := 0;
	vln_luas_bumi_beban_t numeric := 0;
	vln_luas_bng_beban_t numeric := 0;
	vln_njop_bumi_beban_t numeric := 0;
	vln_njop_bng_beban_t numeric := 0;
	vln_flag_op_bersama numeric := 0;
	vln_luas_bumi numeric := 0;
	vln_luas_bng numeric := 0;
	vln_luas_bumi_beban numeric := 0;
	vln_luas_bng_beban numeric := 0;
	vln_njop_bumi numeric := 0;
	vln_njop_bng numeric := 0;
	vln_njop_bumi_beban numeric := 0;
	vln_njop_bng_beban numeric := 0;
	vlc_kd_kls_tanah varchar(3);
	vlc_thn_awal_kls_tanah varchar(4);
	vlc_kd_kls_bng varchar(3);
	vlc_thn_awal_kls_bng varchar(4);
	vlc_kd_kls_tanah_bbn varchar(3);
	vlc_thn_awal_tanah_bbn varchar(4);
	vlc_kd_kls_bng_bbn varchar(3);
	vlc_thn_awal_bng_bbn varchar(4);
	vlc_nm_wp varchar(30);
	vlc_jalan_wp varchar(30);
	vlc_blok_kav_wp varchar(15);
	vlc_rw_wp varchar(2);
	vlc_rt_wp varchar(3);
	vlc_kelurahan_wp varchar(2);
	vlc_kota_wp varchar(2);
	vlc_kd_pos_wp varchar(5);
	vlc_npwp varchar(12);
	vlc_status_pekerjaan_wp varchar(1);
	vln_flag_njoptkp       number;
	vlc_jns_njoptkp varchar(5);
	vln_indeks_range varchar(5);
	vlc_kd_jns_range varchar(5);
	vlc_kd_sk_njop_njkp varchar(5);
	vln_kena_njoptkp       number;
	vln_njoptkp numeric := 0;
	vln_njop_dasar numeric := 0;
	vlc_jns_bumi char(1);
	vlc_kd_jpb char(2);
	vlc_kd_jpb_jpt char(2);
	vln_persen_njkp numeric := 0;
	vln_nilai_njkp numeric := 0;
	vln_nilai_tarif numeric := 1;
	vln_pbb_terhutang numeric := 0;
	vln_pct_pengurangan_jpb numeric := 0;
	vln_nilai_pengurang_jpb numeric := 0;
	vln_pct_pengurangan_permanen numeric := 0;
	vln_nilai_pengurang_permanen numeric := 0;
	vln_pct_pengurangan_pst numeric := 0;
	vln_nilai_pengurang_pst numeric := 0;
	vln_nilai_pengurang_kom numeric := 0;
	vln_faktor_pengurang_pbb numeric := 0;
	vln_pbb_yang_dibayar numeric := 0;
	vln_siklus_sppt numeric := 0;
	vln_status_pembayaran varchar(1);
	vlc_status_tagihan varchar(1);
	vlc_status_cetak varchar(1);
	vln_update_sppt        number;
	vln_njop_kena_pbb numeric := 0;
	vln_njop_perm2	numeric := 0;
	vln_bagi_kelas_bumi numeric := 0;
	vln_bagi_kelas_bng numeric := 0;
	vln_pbb_minimal	numeric := 0;
  vln_pbb_bayar_select numeric := 0;
  current_year varchar(4);

  rec_komp record;
  cursor_komp cursor for 
    SELECT "NilaiYangDikompensasi"
	  FROM  "PenerimaKompensasi"
	  WHERE "Provinsi_Kode_Pemohon"  = substr(plc_nop,1,2)
		  AND "Daerah_Kode_Pemohon"    = substr(plc_nop,3,2)
		  AND "Kecamatan_Kode_Pemohon" = substr(plc_nop,5,3)
      AND "Kelurahan_Kode_Pemohon" = substr(plc_nop,8,3)
		  AND "Blok_Kode_Pemohon"      = substr(plc_nop,11,3)
		  AND "NoUrut_Pemohon"         = substr(plc_nop,14,4)
		  AND "JenisOp_Pemohon"        = substr(plc_nop,18,1)
	    AND "TahunPajakKompensasi"   = plc_tahun;

begin

  current_year := SELECT date_part('year', now());
  NJOPTKP_Individu(substr(plc_nop,1,2),
    substr(plc_nop,3,2),
    substr(plc_nop,5,3),
    substr(plc_nop,8,3),
    substr(plc_nop,11,3),
    substr(plc_nop,14,4),
    substr(plc_nop,18,1),
    current_year);

  ----------------------------------------------------------------------------------
  --- cat : pln_param_sekarang =0 artinya tahun sekarang, pln_param_sekarang = 1 ---
  ---      artinya tahun dahulu                                                  ---
  ----------------------------------------------------------------------------------
  IF pln_param_sekarang = 1 THEN
    vln_luas_bumi       := pln_luas_bumi;
    vln_luas_bng        := pln_luas_bng;
    vln_luas_bumi_beban := pln_luas_bumi_beban;
    vln_luas_bng_beban  := pln_luas_bng_beban;
    vln_njop_bumi       := pln_njop_bumi;
    vln_njop_bng        := pln_njop_bng;
    vln_njop_bumi_beban := pln_njop_bumi_beban;
    vln_njop_bng_beban  := pln_njop_bng_beban;

    -------------------------------------
	  --- Ambil data dari ObjekPajakPbb ---
	  -------------------------------------
    BEGIN
		  SELECT "WajibPajakPbb_Id", "NoPersil"
		  INTO vlc_subjek_pajak_id, vlc_no_persil
      FROM "ObjekPajakPbb"
		  WHERE "Provinsi_Kode"  = substr(plc_nop,1,2)
        AND "Daerah_Kode"    = substr(plc_nop,3,2)
        AND "Kecamatan_Kode" = substr(plc_nop,5,3)
        AND "Kelurahan_Kode" = substr(plc_nop,8,3)
        AND "Blok_Kode"      = substr(plc_nop,11,3)
        AND "NoUrut"         = substr(plc_nop,14,4)
        AND "JenisOp"        = substr(plc_nop,18,1);
      EXCEPTION
        WHEN OTHERS THEN RAISE_APPLICATION_ERROR(-20099, 'Error saat ambil nilai dari ObjekPajakPbb');
	  END;

    ------------------------------------------
	  --- Ambil Jenis Tanah dari dat_op_bumi ---
	  ------------------------------------------
    BEGIN
      SELECT "JenisBumi"
      INTO vlc_jns_bumi
      FROM "ObjekPajakBumi"
      WHERE "Provinsi_Kode" = substr(plc_nop,1,2)
        AND "Daerah_Kode" = substr(plc_nop,3,2)
        AND "Kecamatan_Kode" = substr(plc_nop,5,3)
        AND "Kelurahan_Kode" = substr(plc_nop,8,3)
        AND "Blok_Kode" = substr(plc_nop,11,3)
        AND "NoUrut" = substr(plc_nop,14,4)
        AND "JenisOp" = substr(plc_nop,18,1);
      EXCEPTION
        WHEN OTHERS THEN vlc_jns_bumi := '1';
	  END;

    ----------------------------------------------------------------------
	  --- Ambil JPB Bangunan untuk jaga-jaga apabila diperlukan JPB untuk
	  --- menentukan NJKP
	  ----------------------------------------------------------------------
	  BEGIN
      SELECT "Jpb_Kode"
      INTO   vlc_kd_jpb
      FROM   "ObjekPajakBangunan"
      WHERE  "Provinsi_Kode"  = substr(plc_nop,1,2)
        AND  "Daerah_Kode" = substr(plc_nop,3,2)
        AND  "Kecamatan_Kode" = substr(plc_nop,5,3)
        AND  "Kelurahan_Kode" = substr(plc_nop,8,3)
        AND  "Blok_Kode" = substr(plc_nop,11,3)
        AND  "NoUrut" = substr(plc_nop,14,4)
        AND  "JenisOp" = substr(plc_nop,18,1)
        AND  "NoBangunan" = 1;

      vlc_kd_jpb_jpt := vlc_kd_jpb;
      EXCEPTION
        WHEN OTHERS THEN vlc_kd_jpb_jpt := vlc_jns_bumi;
	  END;

  ELSIF pln_param_sekarang = 0 THEN
	  ----------------------------------------------------------------------
	  --- Ambil data dari ObjekPajakPbb
	  ----------------------------------------------------------------------
	  BEGIN
		  SELECT "TotalLuasBumi", "TotalLuasBangunan", "NjopBumi", "NjopBangunan", "WajibPajakPbb_Id", "NoPersil"
		  INTO vln_temp_luas_bumi, vln_temp_luas_bng, vln_temp_njop_bumi, vln_temp_njop_bng, vlc_subjek_pajak_id, vlc_no_persil
      FROM "ObjekPajakPbb"
		  WHERE "Provinsi_Kode"  = substr(plc_nop,1,2)
        AND "Daerah_Kode"    = substr(plc_nop,3,2)
        AND "Kecamatan_Kode" = substr(plc_nop,5,3)
        AND "Kelurahan_Kode" = substr(plc_nop,8,3)
        AND "Blok_Kode" = substr(plc_nop,11,3)
        AND "NoUrut" = substr(plc_nop,14,4)
        AND "JenisOp" = substr(plc_nop,18,1);
      EXCEPTION
        WHEN OTHERS THEN RAISE_APPLICATION_ERROR(-20100, 'Error saat ambil nilai dari ObjekPajakPbb');
	  END;

    ----------------------------------------------------------------------
	  --- Ambil Jenis Tanah dari dat_op_bumi
	  ----------------------------------------------------------------------
	  BEGIN
      SELECT "JenisBumi"
      INTO vlc_jns_bumi
      FROM "ObjekPajakBumi"
      WHERE "Provinsi_Kode"  = substr(plc_nop,1,2)
        AND "Daerah_Kode"    = substr(plc_nop,3,2)
        AND "Kecamatan_Kode" = substr(plc_nop,5,3)
        AND "Kelurahan_Kode" = substr(plc_nop,8,3)
        AND "Blok_Kode" = substr(plc_nop,11,3)
        AND "NoUrut" = substr(plc_nop,14,4)
        AND "JenisOp" = substr(plc_nop,18,1);
      EXCEPTION
        WHEN OTHERS THEN vlc_jns_bumi := '1';
	  END;

	  ----------------------------------------------------------------------
	  --- Ambil JPB Bangunan untuk jaga-jaga apabila diperlukan JPB untuk
	  --- menentukan NJKP
	  ----------------------------------------------------------------------
	  BEGIN
      SELECT "Jpb_Kode"
      INTO   vlc_kd_jpb
      FROM   "ObjekPajakBangunan"
      WHERE  "Provinsi_Kode"  = substr(plc_nop,1,2)
        AND "Daerah_Kode" = substr(plc_nop,3,2)
        AND "Kecamatan_Kode" = substr(plc_nop,5,3)
        AND "Kelurahan_Kode" = substr(plc_nop,8,3)
        AND "Blok_Kode" = substr(plc_nop,11,3)
        AND "NoUrut" = substr(plc_nop,14,4)
        AND "JenisOp" = substr(plc_nop,18,1)
        AND "NoBangunan" = 1;

      vlc_kd_jpb_jpt := vlc_kd_jpb;
      EXCEPTION
        WHEN OTHERS THEN vlc_kd_jpb_jpt := vlc_jns_bumi;
	  END;

    ----------------------------------------------------------------------
	  --- Cek apakah anggota op bersama
	  ----------------------------------------------------------------------
	  BEGIN
      SELECT "LuasBumiBeban", "LuasBangunanBeban", "NjopBumiBeban", "NjopBangunanBeban"
		  INTO vln_luas_bumi_beban_t, vln_luas_bng_beban_t, vln_njop_bumi_beban_t, vln_njop_bng_beban_t
		  FROM "AnggotaObjekPajak"
		  WHERE "Provinsi_Kode"  = substr(plc_nop,1,2)
        AND "Daerah_Kode"    = substr(plc_nop,3,2)
        AND "Kecamatan_Kode" = substr(plc_nop,5,3)
        AND "Kelurahan_Kode" = substr(plc_nop,8,3)
        AND "Blok_Kode"      = substr(plc_nop,11,3)
        AND "NoUrut"         = substr(plc_nop,14,4)
        AND "JenisOp"        = substr(plc_nop,18,1);

			IF SQL%FOUND THEN
			  vln_flag_op_bersama := 1;
			ELSE
			  vln_luas_bumi_beban_t := 0;
        vln_luas_bng_beban_t  := 0;
        vln_njop_bumi_beban_t := 0;
        vln_njop_bng_beban_t  := 0;
        vln_flag_op_bersama   := 0;
			END IF;
	  EXCEPTION
		  WHEN OTHERS THEN
        vln_luas_bumi_beban_t := 0;
        vln_luas_bng_beban_t  := 0;
        vln_njop_bumi_beban_t := 0;
        vln_njop_bng_beban_t  := 0;
        vln_flag_op_bersama   := 0;
	  END;

    vln_luas_bumi       := vln_temp_luas_bumi;
    vln_luas_bng        := vln_temp_luas_bng;
    vln_luas_bumi_beban := vln_luas_bumi_beban_t;
    vln_luas_bng_beban  := vln_luas_bng_beban_t;
    vln_njop_bumi       := vln_temp_njop_bumi;
    vln_njop_bng        := vln_temp_njop_bng;
    vln_njop_bumi_beban := vln_njop_bumi_beban_t;
    vln_njop_bng_beban  := vln_njop_bng_beban_t;
  ELSE
     RAISE_APPLICATION_ERROR(-20101, 'Error memasukkan parameter tahun');
  END IF;

  ----------------------------------------------------------------------
  --- Ambil data wajib pajak di dat subjek pajak
  ----------------------------------------------------------------------
  BEGIN
	SELECT "Nama_WP", "Jalan_WP", "Blok_Kav_No", "RW", "RT", "Kelurahan_Kode", "Kota_Kode", "Kode_Pos", "NPWP", "StatusPekerjaan"
	INTO vlc_nm_wp, vlc_jalan_wp, vlc_blok_kav_wp, vlc_rw_wp, vlc_rt_wp, vlc_kelurahan_wp, vlc_kota_wp, vlc_kd_pos_wp, vlc_npwp, vlc_status_pekerjaan_wp
	FROM "SubjekPajak"
	WHERE "Id" = vlc_subjek_pajak_id;
  EXCEPTION
    WHEN OTHERS THEN
      RAISE_APPLICATION_ERROR(-20102, 'Error saat ambil nilai dari SubjekPajak');
  END;

  ----------------------------------------------------------------------
  -- Ambil Kelas Tanah
  -- Tahun pajak untuk mengambil kelas tanah
  -- selalu ambil kelas untuk tahun terakhir
  -- 30 Mei 2001 - Pak Edy Sukarno
  ----------------------------------------------------------------------
  vln_bagi_kelas_bumi := 1;
  vln_bagi_kelas_bumi := (vln_njop_bumi/vln_luas_bumi)

  Penentuan_Kelas(1, current_year, vln_bagi_kelas_bumi, vlc_kd_kls_tanah, vlc_thn_awal_kls_tanah, vln_njop_perm2);

  ----------------------------------------------------------------------
  -- Ambil Kelas Bangunan
  -- Tahun pajak untuk mengambil kelas tanah
  -- selalu ambil kelas untuk tahun terakhir
  -- 30 Mei 2001 - Pak Edy Sukarno
  ----------------------------------------------------------------------
  vln_bagi_kelas_bng := 1;
  vln_bagi_kelas_bng := (vln_njop_bng/vln_luas_bng)

  Penentuan_Kelas(2, current_year, vln_bagi_kelas_bng, vlc_kd_kls_bng, vlc_thn_awal_kls_bng, vln_njop_perm2);
  
  ----------------------------------------------------------------------
  -- Ambil Kelas Tanah dan Bangunan OP Bersama
  -- Tahun pajak untuk mengambil kelas tanah dan bangunan
  -- selalu ambil kelas untuk tahun terakhir
  -- 30 Mei 2001 - Pak Edy Sukarno
  ----------------------------------------------------------------------
  IF vln_flag_op_bersama = 1 THEN
    ----------------------------------------------------------------------
    --- Ambil Kelas Tanah OP Bersama
    ----------------------------------------------------------------------
    vln_bagi_kelas_bumi := 1;
    vln_bagi_kelas_bumi := (vln_njop_bumi_beban/vln_luas_bumi_beban)

	  Penentuan_Kelas(1, current_year, vln_bagi_kelas_bumi, vlc_kd_kls_tanah_bbn, vlc_thn_awal_tanah_bbn, vln_njop_perm2);

    ----------------------------------------------------------------------
    --- Ambil Kelas Bangunan
    ----------------------------------------------------------------------
    vln_bagi_kelas_bng := 1;
    vln_bagi_kelas_bng := (vln_njop_bng_beban/vln_luas_bng_beban)

	  Penentuan_Kelas(2, current_year, vln_bagi_kelas_bng, vlc_kd_kls_bng_bbn, vlc_thn_awal_bng_bbn, vln_njop_perm2);
  END IF;

  ----------------------------------------------------------------------
  --- Tentukan apakah NJOPTKP / BTKP
  ----------------------------------------------------------------------
  IF (CAST(plc_tahun AS int) <= 1994) THEN
    --- artinya pakai BTKP
    vln_flag_njoptkp := 0;
    vlc_jns_njoptkp  := '0';
  ELSE
    --- artinya pakai NJOPTKP
    vln_flag_njoptkp := 1;
    vlc_jns_njoptkp  := '1';
  END IF;

  ----------------------------------------------------------------------
  --- Hitung BTKP atau NJOPTKP, diubah oleh : Teguh, tgl : 12/11/2000 --
  --- #tanya --
  ----------------------------------------------------------------------
  BEGIN
    SELECT b."NilaiNJOPTKP"
    INTO   vln_njoptkp
    FROM   "RefThnNJKPNJOPTKPTarif" AS a, "NJOPTKP" AS b
    WHERE  a."JNS_Range_Kode" 		= '0'
      AND  CAST(a."RangeThnAwal" AS int) <= CAST(plc_tahun AS int)
      AND  CAST(a."RangeThnAkhir" AS int) >= CAST(plc_tahun AS int)
      -- AND  b.kd_jns_range 		= a."JNS_Range_Kode"
      -- AND  b.indeks_range 		= a."IndeksRange"
      AND  b."Provinsi_Kode" 		= substr(plc_nop, 1,2)
      AND  b."Daerah_Kode" 			= substr(plc_nop, 3,2);
    EXCEPTION
      WHEN OTHERS THEN vln_njoptkp := 0;
  END;

  vln_njoptkp := vln_njoptkp * 1000;

  ----------------------------------------------------------------------
  --- Menentukan NJOP dasar pengenaan
  ----------------------------------------------------------------------
  IF vln_flag_njoptkp = 1 THEN
	  BEGIN
	    SELECT COUNT(*)
      INTO   vln_kena_njoptkp
      FROM   "SubjekPajakNJOPTKP"
      WHERE  "Id" = rpad(vlc_subjek_pajak_id,30)
        AND  "Provinsi_Kode"  = substr(plc_nop,1,2)
        AND  "Daerah_Kode"    = substr(plc_nop,3,2)
        AND  "Kecamatan_Kode" = substr(plc_nop,5,3)
        AND  "Kelurahan_Kode" = substr(plc_nop,8,3)
        AND  "Blok_Kode"      = substr(plc_nop,11,3)
        AND  "NoUrut"         = substr(plc_nop,14,4)
        AND  "JenisOp"        = substr(plc_nop,18,1);
      EXCEPTION
        WHEN OTHERS THEN vln_kena_njoptkp := 0;
	  END;
  END IF;

  IF vln_flag_njoptkp = 1 THEN
    IF vln_kena_njoptkp = 0 THEN
      vln_njoptkp := 0;
	  END IF;
    vln_njop_dasar := (vln_njop_bumi + vln_njop_bumi_beban + vln_njop_bng + vln_njop_bng_beban) - vln_njoptkp;
  ELSIF vln_flag_njoptkp = 0 THEN
    vln_njop_bng_btkp := ((vln_njop_bng + vln_njop_bng_beban) - vln_njoptkp);
    IF vln_njop_bng_btkp <= 0 THEN
      vln_njop_bng_btkp := 0;
    END IF;

    ---------------------------------------
    --- Diubah TEGUH [22-02-2001]
    --- vln_njop_dasar := (vln_njop_bumi + vln_njop_bumi_beban) - vln_njop_bng_btkp;
    ---------------------------------------
    vln_njop_dasar := (vln_njop_bumi + vln_njop_bumi_beban) + vln_njop_bng_btkp;
  ELSE
    RAISE_APPLICATION_ERROR(-20103,'error pada saat penentuan NJOP dasar pengenaan !');
  END IF;

  IF vln_njop_dasar <= 0 THEN
    vln_njop_dasar := 0;
  END IF;

  vln_njop_kena_pbb := vln_njop_bumi + vln_njop_bumi_beban + vln_njop_bng + vln_njop_bng_beban;

  ----------------------------------------------------------------------
  --- Hitung NJKP
  ----------------------------------------------------------------------
  BEGIN
    SELECT "IndeksRange", "JNS_Range_Kode", "SK_NJOP_NJKP_Kode"
	  INTO 	vln_indeks_range, vlc_kd_jns_range, vlc_kd_sk_njop_njkp
    FROM   "RefThnNJKPNJOPTKPTarif"
    WHERE  CAST("RangeThnAwal" AS int)  <= CAST(plc_tahun AS int)
      AND	 CAST("RangeThnAkhir" AS int) >= CAST(plc_tahun AS int)
      AND  "JNS_Range_Kode"     = '1';
    EXCEPTION
      WHEN OTHERS THEN
      RAISE_APPLICATION_ERROR(-20104, 'Error saat ambil indeks range, etc');
  END;

  IF vlc_kd_jpb IS NULL THEN
    IF vlc_jns_bumi = '1' THEN
      vlc_kd_jpb_jpt := '50';
    ELSE
      vlc_kd_jpb_jpt := '00';
    END IF;
  ELSE
    vlc_kd_jpb_jpt := vlc_kd_jpb;
  END IF;

  BEGIN
 	  SELECT "NilaiNJKP"
	  INTO   vln_persen_njkp
	  FROM   "NJKP"
	  where "IndeksRange" = vln_indeks_range
      AND "JNS_Range_Kode" = vlc_kd_jns_range
      AND "Provinsi_Kode"  = substr(plc_nop, 1,2)
      AND "Daerah_Kode"     = substr(plc_nop, 3,2)
      AND "JPB_JPT_Kode"   = vlc_kd_jpb_jpt
      AND "NJOPMin"    <= vln_njop_dasar
      AND "NJOPMax"    >= vln_njop_dasar;
    EXCEPTION
      WHEN NO_DATA_FOUND THEN
        RAISE_APPLICATION_ERROR(-20105, 'Data nilai persen njkp tidak ditemukan');
      WHEN OTHERS THEN
        RAISE_APPLICATION_ERROR(-20106, 'Error saat ambil nilai persen njkp');
	END;

  -----------------------------------------------------------------------------------
  -- tambahan TEGUH[22-02-2001], merujuk PP 12 th. 1994 dan program recital [P. Edy]
  -- njkp 40% untuk status pekerjaan wp selain 5, indeks = 2,
  -- dan njop > 1 M {sementara, konfirmasi dulu dg. P Dedi njop yang mana}
  -----------------------------------------------------------------------------------
  IF vlc_status_pekerjaan_wp != '5' THEN
    IF vln_njop_dasar >= 1000000000 THEN
      IF CAST(plc_tahun AS int) >= 1994 and CAST(plc_tahun AS int) <= 2000 THEN
        vln_persen_njkp := 20;
      END IF;
    END IF;
  END IF;

	vln_nilai_njkp := ROUND((vln_persen_njkp / 100) * vln_njop_dasar);

  ----------------------------------------------------------------------
	--- Hitung Tarif
  ----------------------------------------------------------------------
	BEGIN
    SELECT "IndeksRange", "JNS_Range_Kode", "SK_NJOP_NJKP_Kode"
    INTO vln_indeks_range, vlc_kd_jns_range, vlc_kd_sk_njop_njkp
    FROM   "RefThnNJKPNJOPTKPTarif"
    WHERE CAST("RangeThnAwal" AS int)  <= CAST(plc_tahun AS int)
      AND	CAST("RangeThnAkhir" AS int)  >= CAST(plc_tahun AS int)
      AND "JNS_Range_Kode"     = '2';
    EXCEPTION
      WHEN OTHERS THEN
        RAISE_APPLICATION_ERROR(-20107, 'Error saat ambil indeks range, etc');
  END;

  -- #tanya
  -- BEGIN
  --   SELECT "NilaiTarif"
	--   into vln_nilai_tarif
	--   FROM "Tarif"
	--   WHERE indeks_range = vln_indeks_range
  --     AND kd_jns_range = vlc_kd_jns_range;
  --   EXCEPTION
	--   WHEN OTHERS THEN
  --     RAISE_APPLICATION_ERROR(-20108, 'Error saat ambil nilai tarif');
	-- END;

	vln_pbb_terhutang := ROUND((vln_nilai_tarif / 100) * vln_nilai_njkp );

  ----------------------------------------------------------------------
  --- Hitung Faktor Pengurang PBB
	--- Cari dulu pengurangan pengenaan JPB
  ----------------------------------------------------------------------
	BEGIN
    SELECT "PCTPenguranganJPB"
	  INTO vln_pct_pengurangan_jpb
	  FROM "PenguranganPengenaanJPB"
	  WHERE "Provinsi_Kode_Pemohon" = substr(plc_nop,1,2)
		  AND "Daerah_Kode_Pemohon"  = substr(plc_nop,3,2)
		  AND "Kecamatan_Kode_Pemohon" = substr(plc_nop,5,3)
      AND "Kelurahan_Kode_Pemohon" = substr(plc_nop,8,3)
		  AND "Blok_Kode_Pemohon" = substr(plc_nop,11,3)
		  AND "NoUrut_Pemohon" = substr(plc_nop,14,4)
		  AND "JenisOp_Pemohon" = substr(plc_nop,18,1)
		  AND "ThnPengenaan" = plc_tahun;
    EXCEPTION
      WHEN OTHERS THEN vln_pct_pengurangan_jpb := 0;
	END;

	vln_nilai_pengurang_jpb := ROUND(vln_pct_pengurangan_jpb/100) * vln_pbb_terhutang;

  ----------------------------------------------------------------------
	--- Cari dulu pengurangan permanen
  ----------------------------------------------------------------------
	BEGIN
	  SELECT "PCTPenguranganPermanen"
	  INTO   vln_pct_pengurangan_permanen
	  FROM   "PenguranganPermanen"
	  WHERE  "Provinsi_Kode_Pemohon"  = substr(plc_nop,1,2)
      AND  "Daerah_Kode_Pemohon" = substr(plc_nop,3,2)
      AND  "Kecamatan_Kode_Pemohon" = substr(plc_nop,5,3)
      AND  "Kelurahan_Kode_Pemohon" = substr(plc_nop,8,3)
      AND  "Blok_Kode_Pemohon" = substr(plc_nop,11,3)
      AND  "NoUrut_Pemohon" = substr(plc_nop,14,4)
      AND  "JenisOp_Pemohon" = substr(plc_nop,18,1)
      AND  "ThnAwal" <= plc_tahun
      AND  "ThnAkhir" >= plc_tahun;
    EXCEPTION
      WHEN OTHERS THEN vln_pct_pengurangan_permanen := 0;
	END;
  
  vln_nilai_pengurang_permanen :=	ROUND((vln_pct_pengurangan_permanen / 100) * (vln_pbb_terhutang - vln_nilai_pengurang_jpb));

  ----------------------------------------------------------------------
  --- Cari dulu pengurangan pst
  ----------------------------------------------------------------------
	BEGIN
	  SELECT "PCTPenguranganPST"
	  INTO   vln_pct_pengurangan_pst
	  FROM  "PenguranganPST"
	  WHERE "Provinsi_Kode_Pemohon"  = substr(plc_nop,1,2)
		  AND "Daerah_Kode_Pemohon" = substr(plc_nop,3,2)
		  AND "Kecamatan_Kode_Pemohon" = substr(plc_nop,5,3)
      AND "Kelurahan_Kode_Pemohon" = substr(plc_nop,8,3)
		  AND "Blok_Kode_Pemohon" = substr(plc_nop,11,3)
		  AND "NoUrut_Pemohon" = substr(plc_nop,14,4)
		  AND "JenisOp_Pemohon" = substr(plc_nop,18,1)
	    AND "ThnPengPST" = plc_tahun;

    EXCEPTION
      WHEN OTHERS THEN vln_pct_pengurangan_pst := 0;
	END;

	vln_nilai_pengurang_pst := ROUND((vln_pct_pengurangan_pst / 100) * (vln_pbb_terhutang - vln_nilai_pengurang_jpb));

  ----------------------------------------------------------------------
	--- Cari dulu penerima kompensasi
  ----------------------------------------------------------------------
	  vln_nilai_kom numeric := 0;
	  vln_nilai_kom_gabung numeric := 0;
	
  BEGIN
    open cursor_komp;
    loop
      fetch cursor_komp into rec_komp;
      -- exit when no more row to fetch
      exit when not found;

      BEGIN
        vln_nilai_kom := rec_komp."NilaiYangDikompensasi";
        vln_nilai_kom_gabung := vln_nilai_kom_gabung + vln_nilai_kom;
      END;

	  end loop;
    close cursor_list_nop;

	  vln_nilai_pengurang_kom := vln_nilai_kom_gabung;
	END;

	vln_faktor_pengurang_pbb := vln_pct_pengurangan_jpb + vln_nilai_pengurang_permanen + vln_nilai_pengurang_pst + vln_nilai_pengurang_kom;
	vln_pbb_yang_dibayar     := vln_pbb_terhutang - vln_faktor_pengurang_pbb;

  ----------------------------------------------------------------------
	--- Cari Ketetapan PBB Minimal [TEGUH, 18-12-2000]
	--- Masukan dari P. Dedi dan P. Edi
  ----------------------------------------------------------------------
	BEGIN
    SELECT "PBBMinimal"
	  INTO   vln_pbb_minimal
	  FROM   "ParameterSPPTSTTSDHKP";
    EXCEPTION
      WHEN OTHERS THEN vln_pbb_minimal := 0;
	END;

	if vln_pbb_yang_dibayar < vln_pbb_minimal then
    vln_pbb_yang_dibayar := vln_pbb_minimal;
	end if;

  ----------------------------------------------------------------------
	--- Insert atau update ke SPPT
  ----------------------------------------------------------------------
	BEGIN
	  SELECT "Siklus_sppt", "StatusPembayaran_sppt", "StatusTagihan_sppt", "StatusCetak_sppt", "PBBygHarusDibayar_sppt"
	  INTO vln_siklus_sppt, vln_status_pembayaran, vlc_status_tagihan, vlc_status_cetak, vln_pbb_bayar_select
	  FROM "Sppt"
	  WHERE "Propinsi_Id"   = substr(plc_nop,1,2)
	  AND   "Dati2_Id"      = substr(plc_nop,3,2)
	  AND   "Kecamatan_Id"  = substr(plc_nop,5,3)
	  AND   "Keluarahan_Id" = substr(plc_nop,8,3)
	  AND   "Blok_Id"       = substr(plc_nop,11,3)
	  AND   "NoUrut"        = substr(plc_nop,14,4)
	  AND   "JenisOP_Id"    = substr(plc_nop,18,1)
	  AND  "TahunPajakskp_sppt" = plc_tahun;

	  IF SQL%FOUND THEN
      vln_update_sppt := 1;
	  ELSE
      vln_update_sppt 	  := 0;
      vln_siklus_sppt 	  := 0;
      vln_status_pembayaran:= '0';
      vlc_status_tagihan   := '0';
		  vlc_status_cetak 	  := '0';
		  vln_pbb_bayar_select := 0;
	  END IF;
    EXCEPTION
      WHEN OTHERS THEN
        vln_update_sppt 	  := 0;
        vln_siklus_sppt 	  := 0;
        vln_status_pembayaran:= '0';
        vlc_status_tagihan   := '0';
        vlc_status_cetak 	  := '0';
        vln_pbb_bayar_select := 0;
	END;

  IF vln_update_sppt > 0 THEN
    --- Edit by Made/Naryo 11-4-2001
    IF (vln_pbb_bayar_select < vln_pbb_yang_dibayar) or vln_status_pembayaran = '2' THEN
      /*Edit by Teguh/Naryo 19-9-2001*/
      --- status bayarnya dikembalikan jadi 0 lagi karena hasil perhitungan lebih
		  --- besar dari pbb yang sudah ada di sppt, berarti dia harus bayar lagi
      vln_status_pembayaran := '0';
    END IF;

    ----------------------------------------------------------------------
    --- UPDATE sppt
    ----------------------------------------------------------------------
    BEGIN
      UPDATE "Sppt"
      SET "Siklus_sppt" = (vln_siklus_sppt+1),
        "KanwilBank_Id" = plc_kd_kanwil_bank,
        "KPPBBbank_Id" = plc_kd_kppbb_bank,
			  "BankTunggal_Id" = plc_bank_tunggal,
        "BankPersepsi_Id" = plc_bank_persepsi,
			  "TP_Id" = plc_bank_tp,
        "NamaWP_sppt" = vlc_nm_wp,
			  "JalanWPskp_sppt" = vlc_jalan_wp,
	   		"BlokKavNoWP_sppt" = vlc_blok_kav_wp,
			  "RwWP_sppt" = vlc_rw_wp,
        "RtWP_sppt" = vlc_rt_wp,
			  "NoPersil_sppt" = vlc_no_persil,
        "KelurahanWP_sppt" = vlc_kelurahan_wp,
			  "KotaWP_sppt" = vlc_kota_wp,
        "PosWPsppt_Id" = vlc_kd_pos_wp,
			  "Npwp_sppt" = vlc_npwp,
        "KelasTanah_Id" = vlc_kd_kls_tanah,
			  "TahunAwalKelasTanah" = vlc_thn_awal_kls_tanah,
        "KelasBangunan_Id" = vlc_kd_kls_bng,
			  "TahunAwalKelasBangunan" = vlc_thn_awal_kls_bng,
        "TanggalJatuhTempo_sppt" = plc_tgl_jatuh_tempo,
			  "LuasBumi_sppt" = vln_luas_bumi,
        "LuasBangunan_sppt" = vln_luas_bng,
			  "NJOPBumi_sppt" = vln_njop_bumi,
        "NJOPBangunan_sppt" = vln_njop_bng,
			  "NJOP_sppt" = vln_njop_kena_pbb,
        "NJOPTKP_sppt" = vln_njoptkp,
			  "NJKPskp_sppt" = vln_persen_njkp,
        "PBBterhutang_sppt" = vln_pbb_terhutang,
        "Faktorpengurangan_sppt" = vln_faktor_pengurang_pbb,
        "PBBygHarusDibayar_sppt" = vln_pbb_yang_dibayar,
			  "StatusPembayaran_sppt" = vln_status_pembayaran,
        "StatusCetak_sppt" = lpad(to_char(to_number(vlc_status_cetak) + 1), 1),
        "TanggalTerbit_sppt" = TO_DATE(plc_tgl_terbit,'DD/MM/YYYY'),
			  "TanggalCetak_sppt" = TO_DATE(plc_tgl_cetak,'DD/MM/YYYY'),
			  "NIPPencetakan_sppt" = plc_nip_pencetak_sppt
	    WHERE "Propinsi_Id"  = substr(plc_nop,1,2)
        AND "Dati2_Id" = substr(plc_nop,3,2)
        AND "Kecamatan_Id" = substr(plc_nop,5,3)
        AND "Keluarahan_Id" = substr(plc_nop,8,3)
        AND "Blok_Id" = substr(plc_nop,11,3)
        AND "NoUrut" = substr(plc_nop,14,4)
        AND "JenisOP_Id" = substr(plc_nop,18,1)
        AND "TahunPajakskp_sppt" = plc_tahun;
        EXCEPTION
          WHEN OTHERS THEN
		      RAISE_APPLICATION_ERROR(-20109, 'Error saat update sppt');
    END;
	ELSE
    BEGIN
      INSERT INTO sppt ("Propinsi_Id", "Dati2_Id",
        "Kecamatan_Id",
        "Keluarahan_Id",
        "Blok_Id",
        "NoUrut",
        "JenisOP_Id",
        "TahunPajakskp_sppt",
        "Siklus_sppt",
        "KanwilBank_Id",
        "KPPBBbank_Id",
        "BankTunggal_Id",
        "BankPersepsi_Id",
        "TP_Id",
        "NamaWP_sppt", 
        "JalanWPskp_sppt", 
        "BlokKavNoWP_sppt",
        "RwWP_sppt",
        "RtWP_sppt",
        "KelurahanWP_sppt",
        "KotaWP_sppt",
        "PosWPsppt_Id",
        "Npwp_sppt",
        "NoPersil_sppt",
        "KelasTanah_Id",
        "TahunAwalKelasTanah",
        "KelasBangunan_Id",
        "TahunAwalKelasBangunan",
        "TanggalJatuhTempo_sppt",
        "LuasBumi_sppt",
        "LuasBangunan_sppt",
        "NJOPBumi_sppt",
        "NJOPBangunan_sppt",
        "NJOP_sppt",
        "NJOPTKP_sppt",
        "NJKPskp_sppt",
        "PBBterhutang_sppt",
        "Faktorpengurangan_sppt",
        "PBBygHarusDibayar_sppt",
        "StatusPembayaran_sppt",
        "StatusTagihan_sppt",
        "StatusCetak_sppt",
        "TanggalTerbit_sppt",
        "TanggalCetak_sppt",
        "NIPPencetakan_sppt")
	    VALUES (substr(plc_nop,1,2),
        substr(plc_nop,3,2),
        substr(plc_nop,5,3),
        substr(plc_nop,8,3),
        substr(plc_nop,11,3),
        substr(plc_nop,14,4),
        substr(plc_nop,18,1),
        plc_tahun,
        1,
        plc_kd_kanwil_bank,
        plc_kd_kppbb_bank,
        plc_bank_tunggal,
        plc_bank_persepsi,
        plc_bank_tp,
        vlc_nm_wp,
        vlc_jalan_wp,
        vlc_blok_kav_wp,
        vlc_rw_wp,
        vlc_rt_wp,
        vlc_kelurahan_wp,
        vlc_kota_wp,
        vlc_kd_pos_wp,
        vlc_npwp,
        vlc_no_persil,
        vlc_kd_kls_tanah,
        vlc_thn_awal_kls_tanah,
        vlc_kd_kls_bng,
        vlc_thn_awal_kls_bng,
        plc_tgl_jatuh_tempo,
        vln_luas_bumi,
        vln_luas_bng,
        vln_njop_bumi,
        vln_njop_bng,
        vln_njop_kena_pbb,
        vln_njoptkp,
        vln_persen_njkp,
        vln_pbb_terhutang,
        vln_faktor_pengurang_pbb,
        vln_pbb_yang_dibayar,
        0,
        '0',
        '0',
        plc_tgl_terbit,
        plc_tgl_cetak,
        plc_nip_pencetak_sppt);
      
      EXCEPTION
        WHEN OTHERS THEN RAISE_APPLICATION_ERROR(-20110, 'Error saat insert sppt');
    END;

	END IF;

  IF vln_flag_op_bersama = 1 THEN
    vln_update_sppt_bersama number;
    BEGIN
      BEGIN
        SELECT COUNT(*)
        INTO   vln_update_sppt_bersama
        FROM   "SpptObjekBersama"
        WHERE  "Propinsi_Id"  = substr(plc_nop,1,2)
        AND    "Dati2_Id"     = substr(plc_nop,3,2)
        AND    "Kecamatan_Id" = substr(plc_nop,5,3)
        AND    "Keluarahan_Id" = substr(plc_nop,8,3)
        AND    "Blok_Id"      = substr(plc_nop,11,3)
        AND    "NoUrut"      = substr(plc_nop,14,4)
        AND    "JenisOP_Id"    = substr(plc_nop,18,1)
        AND    "TahunPajakskp_sppt" = plc_tahun;
	      EXCEPTION
	        WHEN OTHERS THEN vln_update_sppt_bersama := 0;
      END;

      IF vln_update_sppt_bersama > 0 THEN
        BEGIN
          UPDATE "SpptObjekBersama"
            SET "KelasTanah_Id"         = vlc_kd_kls_tanah_bbn,
              "TahunAwalKelasTanah" 	= vlc_thn_awal_tanah_bbn,
              "KelasBangunan_Id" 			= vlc_kd_kls_bng_bbn,
              "TahunAwalKelasBangunan" 	= vlc_thn_awal_bng_bbn,
              "LuasBumiBeban_sppt" = vln_luas_bumi_beban,
              "LuasBangunanBeban_sppt"  = vln_luas_bng_beban,
              "NJOPBumiBeban_sppt" = vln_njop_bumi_beban,
              "NJOPBangunanBeban_sppt"  = vln_njop_bng_beban
	    	    WHERE "Propinsi_Id"   = substr(plc_nop,1,2)
              AND "Dati2_Id"      = substr(plc_nop,3,2)
              AND "Kecamatan_Id"  = substr(plc_nop,5,3)
              AND "Keluarahan_Id" = substr(plc_nop,8,3)
              AND "Blok_Id"       = substr(plc_nop,11,3)
              AND "NoUrut"        = substr(plc_nop,14,4)
              AND "JenisOP_Id"    = substr(plc_nop,18,1)
              AND "TahunPajakskp_sppt" = plc_tahun;
            EXCEPTION
              WHEN OTHERS THEN
                  RAISE_APPLICATION_ERROR(-20111, 'Error saat update sppt op bersama');
	    	END;
      ELSE
        BEGIN
		 	    INSERT INTO "SpptObjekBersama" ("Propinsi_Id", "Dati2_Id", "Kecamatan_Id", "Keluarahan_Id",
              "Blok_Id", "NoUrut", "JenisOP_Id", "TahunPajakskp_sppt",
              "KelasTanah_Id",
              "TahunAwalKelasTanah",
              "KelasBangunan_Id",
              "TahunAwalKelasBangunan",
              "LuasBumiBeban_sppt",
              "LuasBangunanBeban_sppt",
              "NJOPBumiBeban_sppt",
              "NJOPBangunanBeban_sppt")
		 	    VALUES (substr(plc_nop,1,2),
		 		 	  substr(plc_nop,3,2),
				 	  substr(plc_nop,5,3),
            substr(plc_nop,8,3),
				 	  substr(plc_nop,11,3),
				 	  substr(plc_nop,14,4),
				 	  substr(plc_nop,18,1),
            plc_tahun,
				 	  vlc_kd_kls_tanah_bbn,
				 	  vlc_thn_awal_tanah_bbn,
				 	  vlc_kd_kls_bng_bbn,
		 		 	  vlc_thn_awal_bng_bbn,
				 	  vln_luas_bumi_beban,
				 	  vln_luas_bng_beban,
		 		 	  vln_njop_bumi_beban,
				 	  vln_njop_bng_beban);
          EXCEPTION
          WHEN OTHERS THEN
                RAISE_APPLICATION_ERROR(-20112, 'Error saat insert sppt op bersama');
			  END;
      END IF;
    END;
	END IF;

end;$$
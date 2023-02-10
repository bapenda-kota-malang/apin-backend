create or replace function Penetapan_Massal(
  Provinsi_Kode char(2), Daerah_Kode char(2), Kecamatan_Kode char(3), Kelurahan_Kode char(3),
  Blok_Kode char(3), NoUrut char(4), JenisOp char(1), Persil varchar(10),
  NJOPBumi bigint, NJOPBng bigint, LuasBumi bigint, LuasBng bigint,
  wp_id bigint,
  Kanwil_Kode char(2), KPPBB_Kode char(2), Tunggal_Kode char(2), Persepsi_Kode varchar(5), Tp_Kode char(2), 
  NJOPTKP numeric,
  MinB1 numeric, MaxB1 numeric,
  MinB2 numeric, MaxB2 numeric,
  MinB3 numeric, MaxB3 numeric,
  MinB4 numeric, MaxB4 numeric,
  MinB5 numeric, MaxB5 numeric,
  Tahun char(4),
  TjttBk1 date, TjttBk2 date, TjttBk3 date, TjttBk4 date, TjttBk5 date,
  TrbBk1 date, TrbBk2 date, TrbBk3 date, TrbBk4 date, TrbBk5 date,
  Nip varchar(10), NipTgl date
)
returns numeric
language plpgsql    
as $$
declare
  pbb_min numeric := 0;
  tnh_ks char(1);
  komp numeric := 0;
  komp_stm numeric := 0;
  pct_pst numeric := 0;
  pct_prmn numeric := 0;
  pct_jpb numeric := 0;
  njop_1 numeric := 0;
  njop_bumi_beban numeric := 0;
  njop_bng_beban numeric := 0;
  luas_bumi_beban numeric := 0;
  luas_bng_beban numeric := 0;
  njop_bumi_m2 numeric := 0;
  njop_bng_m2 numeric := 0;
  njop_bumi_m2_beban numeric := 0;
  njop_bng_m2_beban numeric := 0;
  kls_bumi char(3);
  thn_bumi char(4);
  nl_m2_bumi numeric := 0;
  kls_bumi_beban char(3);
  thn_bumi_beban char(4);
  nl_m2_bumi_beban numeric := 0;
  kls_bng char(3);
  thn_bng char(4);
  nl_m2_bng numeric := 0;
  kls_bng_beban char(3);
  thn_bng_beban char(4);
  nl_m2_bng_beban numeric := 0;
  njop numeric := 0;
  nm_wp varchar(30);
  jln_wp varchar(30);
  blk_kawp varchar(15);
  rw_wp char(2);
  rt_wp char(3);
  kel_wp varchar(30);
  kota_wp varchar(30);
  kd_pos varchar(5);
  npwp varchar(15);
  flag_njoptkp numeric := 0;
  njoptkp_sppt numeric := 0;
  pbb_tht numeric := 0;
  pbb_tht1 numeric := 0;
  pgrn_jpb numeric := 0;
  pgrn_prmn numeric := 0;
  pgrn_pst numeric := 0;
  prgn_all numeric := 0;
  kompensasi numeric := 0;
  pbb_hrsbyr numeric := 0;
  flag_bersama numeric := 0;
  tjtt date;
  trb date;
  njkp varchar(2);
  tmp numeric := 0;
  temp char(1);
  pbbnol char(1);
  tarif numeric := 0;
  kat numeric := 0;

  rec_komp record;
  cursor_komp cursor for 
    SELECT "NilaiYangDikompensasi"
    FROM "PenerimaKompensasi"
    WHERE "Provinsi_Kode_Kompensasi" = Provinsi_Kode AND
      "Daerah_Kode_Kompensasi" = Daerah_Kode AND
      "Kecamatan_Kode_Kompensasi" = Kecamatan_Kode AND
      "Kelurahan_Kode_Kompensasi" = Kelurahan_Kode AND
      "Blok_Kode_Kompensasi" = Blok_Kode AND
      "NoUrut_Kompensasi" =  NoUrut AND
      "JenisOp_Kompensasi" = JenisOp AND
      "TahunPajakKompensasi" = Tahun;
  
  rec_komp2 record;
  cursor_komp2 cursor for 
    SELECT "NilaiYangDikompensasi"
    FROM "PenerimaKompensasi"
    WHERE "Provinsi_Kode_Kompensasi" = Provinsi_Kode AND
      "Daerah_Kode_Kompensasi" = Daerah_Kode AND
      "Kecamatan_Kode_Kompensasi" = Kecamatan_Kode AND
      "Kelurahan_Kode_Kompensasi" = Kelurahan_Kode AND
      "Blok_Kode_Kompensasi" = Blok_Kode AND
      "NoUrut_Kompensasi" =  NoUrut AND
      "JenisOp_Kompensasi" = JenisOp AND
      "TahunPajakKompensasi" = CAST(Tahun AS int)-1;

  rec_cek_bersama record;
  cursor_cek_bersama cursor for 
    SELECT "LuasBumiBeban", "LuasBangunanBeban", "NjopBumiBeban", "NjopBangunanBeban"
    FROM "AnggotaObjekPajak"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" =  NoUrut AND
      "JenisOp" = JenisOp;
  
  rec_cek_sppt record;
  cursor_cek_sppt cursor for 
    SELECT "Siklus_sppt"
    FROM "Sppt"
    WHERE "Propinsi_Id" = Provinsi_Kode AND
      "Dati2_Id" = Daerah_Kode AND
      "Kecamatan_Id" = Kecamatan_Kode AND
      "Keluarahan_Id" = Kelurahan_Kode AND
      "Blok_Id" = Blok_Kode AND
      "NoUrut" =  NoUrut AND
      "JenisOP_Id" = JenisOp AND
      "TahunPajakskp_sppt" = Tahun;
  
  rec_cek_sppt_bersama record;
  cursor_cek_sppt_bersama cursor for 
    SELECT "Siklus_sppt"
    FROM "Sppt"
    WHERE "Propinsi_Id" = Provinsi_Kode AND
      "Dati2_Id" = Daerah_Kode AND
      "Kecamatan_Id" = Kecamatan_Kode AND
      "Keluarahan_Id" = Kelurahan_Kode AND
      "Blok_Id" = Blok_Kode AND
      "NoUrut" =  NoUrut AND
      "JenisOP_Id" = JenisOp AND
      "TahunPajakskp_sppt" = Tahun;

begin

  -- cek status OP bersama
  open cursor_cek_bersama;

  loop
    fetch cursor_cek_bersama into rec_cek_bersama;
    -- exit when no more row to fetch
    exit when not found;

    -- cari jumlah satuan untuk masing2 kode fasilitas dari tabel "FasilitasBangunan"
    BEGIN
      luas_bumi_beban := rec_cek_bersama."LuasBumiBeban";
      luas_bng_beban := rec_cek_bersama."LuasBangunanBeban";
      njop_bumi_beban := rec_cek_bersama."NjopBumiBeban";
      njop_bng_beban := rec_cek_bersama."NjopBangunanBeban";
	  END;

  end loop;
  
  close cursor_cek_bersama;

  -- cari njoptkp di dat_subjek_pajak_njoptkp
  BEGIN
    SELECT count(*) into flag_njoptkp
    FROM "SubjekPajakNJOPTKP"
    Where "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" =  NoUrut AND
      "JenisOp" = JenisOp;
  
    exception when no_data_found then flag_njoptkp := 0;
  end;

  if flag_njoptkp = 0 then
    njoptkp_sppt := 0;
  else
    njoptkp_sppt := NJOPTKP;
  end if;

  -- total NJOP
  njop_1 := (NJOPBumi + NJOPBng + njop_bumi_beban + njop_bng_beban);
  njop   := (NJOPBumi + NJOPBng + njop_bumi_beban + njop_bng_beban) - njoptkp_sppt;

  if njop < 0 then
    njop := 0;
  end if;

  -- cari tarif
  BEGIN
    SELECT "NilaiTarif" INTO tarif
    FROM "Tarif"
    WHERE "Provinsi_Kode" = Provinsi_Kode
      AND "Daerah_Kode" = Daerah_Kode
      AND "ThnAwal" >= Tahun
      AND "ThnAkhir" <= Tahun
      AND "NJOPMin" >= njop
      AND "NJOPMax" <= njop;
    
    exception when others then tarif:= 100;
  end;
  
  -- PBB terhutang
  pbb_tht := round(njop * (tarif/100));

  -- cari pengurangan JPB
  BEGIN
    BEGIN
      SELECT "PCTPenguranganJPB" INTO pct_jpb
      FROM "PenguranganPengenaanJPB"
      WHERE "Provinsi_Kode_Pemohon" = Provinsi_Kode AND
        "Daerah_Kode_Pemohon" = Daerah_Kode AND
        "Kecamatan_Kode_Pemohon" = Kecamatan_Kode AND
        "Kelurahan_Kode_Pemohon" = Kelurahan_Kode AND
        "Blok_Kode_Pemohon" = Blok_Kode AND
        "NoUrut_Pemohon" = NoUrut AND
        "JenisOp_Pemohon" = JenisOp AND
        "ThnPengenaan" = Tahun;

      exception when others then pct_jpb := 0;
    end;
    
    if npct_jpb > 0 then
      pgrn_jpb := round(pbb_tht * (pct_jpb/100));
    else pgrn_jpb := 0;
    end if;
    
    exception when others then pgrn_jpb := 0;
  END;

  pbb_tht1 := pbb_tht - pgrn_jpb;

  -- cari pengurangan permanen
  Begin
    Begin
      SELECT "PCTPenguranganPermanen" INTO pct_prmn
      from "PenguranganPermanen"
      WHERE "Provinsi_Kode_Pemohon" = Provinsi_Kode AND
        "Daerah_Kode_Pemohon" = Daerah_Kode AND
        "Kecamatan_Kode_Pemohon" = Kecamatan_Kode AND
        "Kelurahan_Kode_Pemohon" = Kelurahan_Kode AND
        "Blok_Kode_Pemohon" = Blok_Kode AND
        "NoUrut_Pemohon" = NoUrut AND
        "JenisOp_Pemohon" = JenisOp AND
        "ThnAwal" >= Tahun AND
        "ThnAkhir" <= Tahun;
      
      exception when others then pct_prmn := 0;
    END;

    if pct_prmn > 0 then
      pgrn_prmn := round(pbb_tht1 * (pct_prmn/100));
    else pgrn_prmn := 0;
    end if;
  
    exception when others then pgrn_prmn := 0;
  END;

  -- cari pengurangan pst
  BEGIN
    BEGIN
      SELECT "PCTPenguranganPST" into pct_pst
      from "PenguranganPST"
      WHERE "Provinsi_Kode_Pemohon" = Provinsi_Kode AND
        "Daerah_Kode_Pemohon" = Daerah_Kode AND
        "Kecamatan_Kode_Pemohon" = Kecamatan_Kode AND
        "Kelurahan_Kode_Pemohon" = Kelurahan_Kode AND
        "Blok_Kode_Pemohon" = Blok_Kode AND
        "NoUrut_Pemohon" = NoUrut AND
        "JenisOp_Pemohon" = JenisOp AND
        "ThnPengPST" = Tahun;
    
      exception when others then pct_pst := 0;
    END;
    
    if pct_pst > 0 then
      pgrn_pst := round(pbb_tht1 * (pct_pst/100));
    else pgrn_pst := 0;
    end if;
  
    exception when others then pgrn_pst := 0;
  END;

  -- cari kompensasi
  kompensasi := 0;
  begin
    open cursor_komp;

    loop
      fetch cursor_komp into rec_komp;
      -- exit when no more row to fetch
      exit when not found;

      kompensasi := rec_komp;
    
    end loop;
  
    close cursor_komp;
  end;

  -- cari kelas tanah
  IF (NJOPBumi <> 0) AND (LuasBumi <> 0) THEN
    BEGIN
      njop_bumi_m2 := (NJOPBumi/LuasBumi)/1000;
      SELECT "KdTanah", "TahunAwalKelasTanah", "NilaiPerM2Tanah"
      INTO kls_bumi, thn_bumi, nl_m2_bumi FROM "KelasTanah"
      WHERE 
        -- thn_awal_kls_tanah <= TO_CHAR(SYSDATE,'yyyy') AND
        -- thn_akhir_kls_tanah >= TO_CHAR(SYSDATE,'yyyy') AND
        "NilaiMinTanah" < njop_bumi_m2 AND
        "NilaiMaxTanah" >= njop_bumi_m2 AND
        "KdTanah" NOT IN ('XXX', '00');
      
      EXCEPTION WHEN OTHERS THEN
        kls_bumi := 'XXX';
        thn_bumi := '1986';
        nl_m2_bumi := njop_bumi_m2;
    END;
  ELSE
    kls_bumi := 'XXX';
    thn_bumi := '1986';
    nl_m2_bumi := 0;
  END IF;

  -- cari kelas bangunan
  IF (NJOPBng <> 0) AND (LuasBng <> 0) THEN
    BEGIN
      njop_bng_m2 := (NJOPBng/LuasBng)/1000;
      SELECT "KdBangunan", "TahunAwalKelasBangunan", NilaiPerM2Bangunan
      INTO kls_bng, thn_bng, nl_m2_bng FROM "KelasBangunan"
      WHERE 
        -- thn_awal_kls_bng <= TO_CHAR(SYSDATE,'yyyy') AND
        -- thn_akhir_kls_bng >= TO_CHAR(SYSDATE,'yyyy') AND
        "NilaiMinBangunan" < njop_bng_m2 AND
        "NilaiMaxBangunan" >= njop_bng_m2 AND
        "KdBangunan" NOT IN ('XXX', '00');
    
      EXCEPTION WHEN OTHERS THEN
        kls_bng := 'XXX';
        thn_bng := '1986';
        nl_m2_bng := njop_bng_m2;
    END;
  ELSE
    kls_bng := 'XXX';
    thn_bng := '1986';
    nl_m2_bng := 0;
  END IF;

  -- cari kelas tanah  beban
  IF flag_bersama = 1 THEN
    BEGIN
      -- cari kelas tanah
      IF (njop_bumi_beban <> 0) AND (luas_bumi_beban <> 0) THEN
        BEGIN
          njop_bumi_m2_beban := (njop_bumi_beban/luas_bumi_beban)/1000;
          SELECT "KdTanah", "TahunAwalKelasTanah", "NilaiPerM2Tanah"
          INTO kls_bumi_beban, thn_bumi_beban, nl_m2_bumi_beban FROM "KelasTanah"
          WHERE
            -- thn_awal_kls_tanah <= TO_CHAR(SYSDATE,'yyyy') AND
            -- thn_akhir_kls_tanah >= TO_CHAR(SYSDATE,'yyyy') AND
            "NilaiMinTanah" < njop_bumi_m2_beban AND
            "NilaiMaxTanah" >= njop_bumi_m2_beban AND
            "KdTanah" NOT IN ('XXX', '00');
        
          EXCEPTION WHEN OTHERS THEN
            kls_bumi_beban := 'XXX';
            thn_bumi_beban := '1986';
            nl_m2_bumi_beban := njop_bumi_m2_beban;
        END;
      ELSE
        kls_bumi_beban := 'XXX';
        thn_bumi_beban := '1986';
        nl_m2_bumi_beban := 0;
      END IF;

      -- cari kelas bangunan
      IF (njop_bng_beban <> 0) AND (luas_bng_beban <> 0) THEN
        BEGIN
          njop_bng_m2_beban := (njop_bng_beban/luas_bng_beban)/1000;
          SELECT "KdBangunan", "TahunAwalKelasBangunan", "NilaiPerM2Bangunan"
          INTO kls_bng_beban, thn_bng_beban, nl_m2_bng_beban FROM "KelasBangunan"
          WHERE
            -- thn_awal_kls_bng <= TO_CHAR(SYSDATE,'yyyy') AND
            -- thn_akhir_kls_bng >= TO_CHAR(SYSDATE,'yyyy') AND
            "NilaiMinBangunan" < njop_bng_m2_beban AND
            "NilaiMaxBangunan" >= njop_bng_m2_beban AND
            "KdBangunan" NOT IN ('XXX', '00');
        
          EXCEPTION WHEN OTHERS THEN
            kls_bng_beban := 'XXX';
            thn_bng_beban := '1986';
            nl_m2_bng_beban := njop_bng_m2_beban;
        END;
      ELSE
        kls_bng_beban := 'XXX';
        thn_bng_beban := '1986';
        nl_m2_bng_beban := 0;

      END IF;
      
    EXCEPTION WHEN OTHERS THEN NULL;
    END;
  END IF;

  -- lookup subjek_pajak
  begin
    Select "Nama_WP", "Jalan_WP", "Blok_Kav_No", "RT", "RW", "Kelurahan_Kode", "Kota_Kode", "Kode_Pos", "NPWP"
    into nm_wp, jln_wp, blk_kawp, rw_wp, rt_wp, kel_wp, kota_wp, kd_pos, npwp
    from "SubjekPajak"
    Where "NPWP" IN (SELECT "NPWP" FROM "Sppt" WHERE "Propinsi_Id" = Provinsi_Kode AND
      "Dati2_Id" = Daerah_Kode AND
      "Kecamatan_Id" = Kecamatan_Kode AND
      "Keluarahan_Id" = Kelurahan_Kode AND
      "Blok_Id" = Blok_Kode AND
      "NoUrut" =  NoUrut AND
      "JenisOP_Id" = JenisOp AND
      "TahunPajakskp_sppt" = Tahun);
  
    exception when others then null;
  end;

  -- pengurangan total
  prgn_all := pgrn_jpb + pgrn_prmn + pgrn_pst + kompensasi;

  --  pbb yh harus dibayar
  pbb_hrsbyr := pbb_tht - prgn_all;

  if pbb_tht between MinB1 and MaxB1 then
    tjtt := TjttBk1;
    trb :=  TrbBk1;
  elsif pbb_tht between MinB2 and MaxB2 then
    tjtt := TjttBk2;
    trb :=  TrbBk2;
  elsif pbb_tht between MinB3 and MaxB3 then
    tjtt := TjttBk3;
    trb :=  TrbBk3;
  elsif pbb_tht between MinB4 and MaxB4 then
    tjtt := TjttBk4;
    trb :=  TrbBk4;
  elsif pbb_tht between MinB5 and MaxB5 then
    tjtt := TjttBk5;
    trb :=  TrbBk5;
  end if;


  if pbb_hrsbyr <= pbb_min then
    if pbb_min > 0 then
      pbbnol := '0';
    else
      pbbnol := '1';
    end if;
    pbb_hrsbyr := pbb_min;
  else pbbnol := '0';
  end if;

  ---CEK APAKAH KATEGORI 3 ATAU TIDAK
  BEGIN
    SELECT count(*) into kat from "DafnomOP"
    where "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" =  NoUrut AND
      "JenisOp" = JenisOp AND
      "Kategori_OP" = '3';

    EXCEPTION WHEN OTHERS THEN
      kat:=0;
  END;

  if kat = 0 then
    -- insert or update ke SPPT
    begin
      open cursor_cek_sppt;
      fetch cursor_cek_sppt into rec_cek_sppt;
      if found then
        Update "Sppt" set "Siklus_sppt" = (rec_cek_sppt."Siklus_sppt") + 1, "KanwilBank_Id" = Kanwil_Kode,
          "KPPBBbank_Id" = KPPBB_Kode, "BankTunggal_Id" = Tunggal_Kode, 
          "BankPersepsi_Id" = Persepsi_Kode, "TP_Id" = Tp_Kode, "NamaWP_sppt" = nm_wp,
          "JalanWPskp_sppt" = jln_wp, "BlokKavNoWP_sppt" = blk_kawp,
          "RwWP_sppt" = rw_wp, "RtWP_sppt" = rt_wp, "KelurahanWP_sppt" = kel_wp,
          "KotaWP_sppt" = kota_wp, "PosWPsppt_Id" = kd_pos, "Npwp_sppt" = npwp,
          "NoPersil_sppt" = persil, "KelasTanah_Id" = kls_bumi,
          "TahunAwalKelasTanah" = thn_bumi, "KelasBangunan_Id" = kls_bng,
          "TahunAwalKelasBangunan" = thn_bng, "TanggalJatuhTempo_sppt" = tjtt,
          "LuasBumi_sppt" = LuasBumi, "LuasBangunan_sppt" = LuasBng,
          "NJOPBumi_sppt" = NJOPBumi, "NJOPBangunan_sppt" = NJOPBng, "NJOP_sppt" = njop_1,
          "NJOPTKP_sppt" = njoptkp_sppt, "NJKPskp_sppt" = 0, "PBBterhutang_sppt" = pbb_tht,
          "Faktorpengurangan_sppt" = prgn_all, "PBBygHarusDibayar_sppt" = pbb_hrsbyr,
          "StatusPembayaran_sppt" = 0, "StatusTagihan_sppt" = '0', "StatusCetak_sppt" = '1',
          "TanggalTerbit_sppt" = trb, "TanggalCetak_sppt" = NipTgl, "NIPPencetakan_sppt" = nip
        where "Propinsi_Id" = Provinsi_Kode AND
          "Dati2_Id" = Daerah_Kode AND
          "Kecamatan_Id" = Kecamatan_Kode AND
          "Keluarahan_Id" = Kelurahan_Kode AND
          "Blok_Id" = Blok_Kode AND
          "NoUrut" =  NoUrut AND
          "JenisOP_Id" = JenisOp AND
          "TahunPajakskp_sppt" = Tahun;
      else
        Insert into sppt("Propinsi_Id", "Dati2_Id", "Kecamatan_Id", "Keluarahan_Id", "Blok_Id", "NoUrut", "JenisOP_Id", "TahunPajakskp_sppt",
          "Siklus_sppt", "KanwilBank_Id", "KPPBBbank_Id", "BankTunggal_Id", "BankPersepsi_Id", "TP_Id",
          "NamaWP_sppt", "JalanWPskp_sppt", "BlokKavNoWP_sppt", "RwWP_sppt", "RtWP_sppt", "KelurahanWP_sppt",
          "KotaWP_sppt", "PosWPsppt_Id", "Npwp_sppt", "NoPersil_sppt", "KelasTanah_Id",
          "TahunAwalKelasTanah", "KelasBangunan_Id", "TahunAwalKelasBangunan", "TanggalJatuhTempo_sppt", "LuasBumi_sppt",
          "LuasBangunan_sppt", "NJOPBumi_sppt", "NJOPBangunan_sppt", "NJOP_sppt", "NJOPTKP_sppt", "NJKPskp_sppt", "PBBterhutang_sppt",
          "Faktorpengurangan_sppt", "PBBygHarusDibayar_sppt", "StatusPembayaran_sppt", "StatusTagihan_sppt",
          "StatusCetak_sppt", "TanggalTerbit_sppt", "TanggalCetak_sppt", "NIPPencetakan_sppt")
        Values (Provinsi_Kode, Daerah_Kode, Kecamatan_Kode, Kelurahan_Kode, Blok_Kode, NoUrut, JenisOp, Tahun, 1, 
          Kanwil_Kode, KPPBB_Kode, Tunggal_Kode, Persepsi_Kode, Tp_Kode,
          nm_wp, jln_wp, blk_kawp, rw_wp, rt_wp, kel_wp,
          kota_wp, kd_pos, npwp, persil, kls_bumi,
          thn_bumi, kls_bng, thn_bng, tjtt, LuasBumi,
          LuasBng, NJOPBumi, NJOPBng, njop_1, njoptkp_sppt, 0, pbb_tht, prgn_all,
          pbb_hrsbyr, pbbnol, '0', '1', trb, NipTgl, nip);
      end if;
      Close cursor_cek_sppt;
    end;

    -- insert ke sppt op anggota
    if flag_bersama = 1 then
      begin
        open cursor_cek_sppt_bersama;
        fetch cursor_cek_sppt_bersama into rec_cek_sppt_bersama;
        if found then
          update SpptObjekBersama set "KelasTanah_Id" = kls_bumi_beban,
            "TahunAwalKelasTanah" = thn_bumi_beban, "KelasBangunan_Id" = kls_bng_beban,
            "TahunAwalKelasBangunan" = thn_bng_beban, "LuasBumiBeban_sppt" = luas_bumi_beban,
            "LuasBangunanBeban_sppt" = luas_bng_beban, "NJOPBumiBeban_sppt" = njop_bumi_beban,
            "NJOPBangunanBeban_sppt" = njop_bng_beban
          where "Propinsi_Id" = Provinsi_Kode AND
          "Dati2_Id" = Daerah_Kode AND
          "Kecamatan_Id" = Kecamatan_Kode AND
          "Keluarahan_Id" = Kelurahan_Kode AND
          "Blok_Id" = Blok_Kode AND
          "NoUrut" =  NoUrut AND
          "JenisOP_Id" = JenisOp AND
          "TahunPajakskp_sppt" = Tahun;
        else
          insert into sppt_op_bersama("Propinsi_Id", "Dati2_Id", "Kecamatan_Id",
            "Keluarahan_Id", "Blok_Id", "NoUrut", "JenisOP_Id", "TahunPajakskp_sppt", "KelasTanah_Id",
            "TahunAwalKelasTanah", "KelasBangunan_Id", "TahunAwalKelasBangunan", "LuasBumiBeban_sppt",
            "LuasBangunanBeban_sppt", "NJOPBumiBeban_sppt", "NJOPBangunanBeban_sppt")
          values (prop, dat, kec, kel, blk, urut, jns,
            Tahun, kls_bumi_beban, thn_bumi_beban, kls_bng_beban,
            thn_bng_beban, luas_bumi_beban, luas_bng_beban, njop_bumi_beban,
            njop_bng_beban);
        end if;
        close cursor_cek_sppt_bersama;
      end;
    end if;

    if pbbnol = '1' then 
      begin
        Update sppt set status_pembayaran_sppt = 1
        where "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Kecamatan_Kode" = Kecamatan_Kode AND
          "Kelurahan_Kode" = Kelurahan_Kode AND
          "Blok_Kode" = Blok_Kode AND
          "NoUrut" =  NoUrut AND
          "JenisOp" = JenisOp AND
          "TahunPajakskp_sppt" = Tahun;
        exception when others then null;
      end;
    end if;
  end if;

  return pbb_min;

end;$$

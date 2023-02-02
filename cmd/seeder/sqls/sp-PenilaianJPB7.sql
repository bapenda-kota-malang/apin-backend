DROP FUNCTION IF EXISTS "public"."penilaian_jpb7"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar);
CREATE OR REPLACE FUNCTION "public"."penilaian_jpb7"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_jpb numeric := 0;
  nilai_jpb7 numeric := 0;
  kls_jpb6 char(1);
	nilai_komponen_utama numeric := 0;
	thn_dibangun_bng char(4);
  thn_renovasi_bng char(4);
  luas_bng numeric := 0;
  kondisi_bng char(1);
	atap char(1);
	lantai char(1);
	langit_langit char(1);
	jml_lantai_bng int := 1;
	nilai_atap numeric := 0;
	nilai_lantai numeric:= 0;
	nilai_langit_langit numeric := 0;
	dinding char(1);
	cari_dinding char(2);
	nilai_dinding_temp numeric := 0;
	nilai_dinding_plester numeric := 0;
	nilai_dinding_cat numeric := 0;
	nilai_dinding numeric := 0;
	nilai_material numeric := 0;
	nilai_fasilitas numeric := 0;
	nilai_total_per_m2 numeric := 0;
	nilai_total_kali_luas numeric := 0;
	nilai_sebelum_susut numeric := 0;
	persentase_penyusutan numeric := 0;
	persen_susut numeric := 0;
	nilai_setelah_susut numeric := 0;
	nilai_fasilitas_susut numeric := 0;
	nilai_fasilitas_tdk_susut numeric := 0;
	penyusutan numeric := 0;
	nilai_dbkb_jpb7 numeric := 0;
  kd_fasilitas char(2);
  status_fasilitas  varchar(4);
  ketergantungan char(2);
  nilai_satuan numeric := 0;
  jns_jpb7 char(1);
  bintang_jpb7 char(1);
  luas_jpb7 numeric := 0;
  jml_kmr_jpb7 numeric := 0;
  luas_jpb7_lain numeric := 0;
  jml_satuan numeric := 0;
  rec_fasilitas record;
  cur_fasilitas cursor for SELECT "Status", "Fasilitas_Kode", "Ketergantungan" FROM "Fasilitas" WHERE "Status" IN ('0','1','2','3');

begin

  -- cari nilai komponen utama dbkb jpb 7 dari tabel DBKB_JPB7
  BEGIN
    SELECT a."TahunDibangun", a."TahunRenovasi", a."LuasBangunan", a."JmlLantaiBng",
      a."Kondisi", b."JenisHotel", b."JumlahBintang", b."JumlahKamar", b."LuasKamarAcCentral", b."LuasRuangLainAcCentral"
    INTO thn_dibangun_bng, thn_renovasi_bng, luas_bng, jml_lantai_bng, kondisi_bng,
      jns_jpb7, bintang_jpb7, jml_kmr_jpb7, luas_jpb7, luas_jpb7_lain
    FROM "ObjekPajakBangunan" AS a, "Jpb7" AS b
    WHERE a."Provinsi_Kode" = Provinsi_Kode
      AND a."Daerah_Kode" = Daerah_Kode
      AND a."Kecamatan_Kode" = Kecamatan_Kode
      AND a."Kelurahan_Kode" = Kelurahan_Kode
      AND a."Blok_Kode" = Blok_Kode
      AND a."NoUrut" = NoUrut
      AND a."JenisOp" = JenisOp
      AND a."NoBangunan" = NoBangunan
      AND a."Jpb_Kode" = '07'
      AND b."Provinsi_Kode" = a."Provinsi_Kode"
      AND b."Daerah_Kode" = a."Daerah_Kode"
      AND b."Kecamatan_Kode" = a."Kecamatan_Kode"
      AND b."Kelurahan_Kode" = a."Kelurahan_Kode"
      AND b."Blok_Kode" = a."Blok_Kode"
      AND b."NoUrut" = a."NoUrut"
      AND b."JenisOp" = a."JenisOp"
      AND b."NoBangunan" = a."NoBangunan";

  END;

  -- nilai komponen utama dbkb jpb 7 dari tabel DBKB_JPB7
  -- jika bintang '0' ubah menjadi bintang '5',
  -- diedit Pak Edy, Tgl. 21/10/2000
  IF bintang_jpb7 = '0' THEN
    bintang_jpb7 := '5';
  END IF;

  BEGIN
    SELECT "NilaiDbkbJp75"
    INTO nilai_dbkb_jpb7
    FROM "DbkbJpb7"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "TahunDbkbJpb7" = Tahun AND
      "JenisDbkbJpb7" = jns_jpb7 AND
      "BintangDbkbJpb7" = bintang_jpb7 AND
      "LantaiMinJpb7" <= jml_lantai_bng AND
      "LantaiMaxJpb7" >= jml_lantai_bng;

    if not found then
      nilai_dbkb_jpb7 = 0;
    end if;
  END;

  -- nilai komponen utama X luas bangunan
  nilai_dbkb_jpb7 := nilai_dbkb_jpb7 * luas_bng;

  -- nilai biaya fasilitas yang dipengaruhi luas, jumlah kamar, luas kamar
  nilai_fasilitas := 0;

  -- proses cur_fasilitas
  open cur_fasilitas;

  loop
    fetch cur_fasilitas into rec_fasilitas;
    -- exit when no more row to fetch
    exit when not found;

    BEGIN
      SELECT "JumlahSatuan" INTO jml_satuan
      FROM "FasilitasBangunan"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
        "Daerah_Kode" = Daerah_Kode AND
        "Kecamatan_Kode" = Kecamatan_Kode AND
        "Kelurahan_Kode" = Kelurahan_Kode AND
        "Blok_Kode" = Blok_Kode AND
        "NoUrut" =  NoUrut AND
        "JenisOp" = JenisOp AND
        "NoBangunan" = NoBangunan AND
        "KodeFasilitas" = rec_biaya_fasilitas."Fasilitas_Kode";
	  
      if not found then
        jml_satuan = 0;
      end if;
	  END;

    IF rec_fasilitas."Ketergantungan" = '0' THEN
      BEGIN
        SELECT "NilaiNonDep"
			  INTO nilai_satuan
        FROM "FasNonDep"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Thn_Non_Dep" = Tahun AND
          "Fasilitas_Kode" = kd_fasilitas;

        if not found then
          nilai_satuan = 0;
        end if;
      END;
	  ELSIF rec_fasilitas."Ketergantungan" = '1' THEN
      BEGIN
        SELECT "NilaiDepMinMax"
        INTO nilai_satuan
        FROM "FasDepMinMax"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "TahunDepMinMax" = Tahun AND
          "Fasilitas_Kode" = kd_fasilitas AND
          "KlsDepMin" <= jml_satuan AND
          "KlsDepMax" >= jml_satuan;

        if not found then
          nilai_satuan = 0;
        end if;
      END;
	  ELSIF rec_fasilitas."Ketergantungan" = '2' THEN
      BEGIN
        SELECT "NilaiFasilitasKlsBintang"
        INTO nilai_satuan
        FROM "FasDepJpbKlsBintang"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "TahunDep" = Tahun AND
          "Fasilitas_Kode" = kd_fasilitas AND
          "Jpb_Kode" = '07' AND
          "KlsBintang" = bintang_jpb7;

        if not found then
          nilai_satuan = 0;
        end if;
      END;
	  ELSE
      nilai_satuan := 0;
	  END IF;

    -- penentuan biaya fasilitas yang dipengaruhi luas, jumlah kamar, luas kamar
    IF status_fasilitas = '0'  THEN
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * jml_satuan * luas_bng);
	  ELSIF status_fasilitas = '1'  THEN
      ---------------------------------------------------------------------
      -- cari nilai boiler, anggap selalu ada
      ---------------------------------------------------------------------
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * jml_kmr_jpb7);
	  ELSIF status_fasilitas = '2'  THEN
      ---------------------------------------------------------------------
      -- cari ac sentral kamar hotel, tergantung luas-nya
      ---------------------------------------------------------------------
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * luas_jpb7);
	  ELSIF status_fasilitas = '3'  THEN
      ---------------------------------------------------------------------
      -- cari ac sentral non kamar hotel, tergantung luas-nya
      ---------------------------------------------------------------------
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * luas_jpb7_lain);
	  ELSE
      nilai_fasilitas := nilai_fasilitas;
	  END IF;
    
  end loop;
  
  close cur_fasilitas;

  -- nilai komponen utama + nilai biaya fasilitas
  nilai_jpb7 := nilai_dbkb_jpb7 + nilai_fasilitas;

  -- nilai biaya fasilitas yang disusutkan
  BEGIN
    SELECT "fasilitas_susut"
    INTO nilai_fasilitas_susut
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

  -- cari nilai sebelum disusutkan
  nilai_jpb7 := nilai_jpb7 + nilai_fasilitas_susut;

  -- cari besar prosentase penyusutan
  BEGIN
    SELECT "susut"
    INTO persentase_penyusutan
    FROM Susut(
      CAST(Tahun AS int),
      CAST(thn_dibangun_bng AS int),
      CAST(thn_renovasi_bng AS int),
      kondisi_bng,
      nilai_jpb7,
      luas_bng,
      0
    );

  END;

  penyusutan := ROUND((penyusutan * nilai_jpb7) / 100);

  -- cari nilai setelah penyusutan
  nilai_jpb7 := nilai_jpb7 - penyusutan;

  -- cari nilai fasilitas yang tidak disusutkan
  BEGIN
    SELECT "fasilitas_tidak_susut"
    INTO nilai_fasilitas_tdk_susut
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

  -- cari nilai bangunan jpb 7
  nilai_jpb := nilai_jpb7 + nilai_fasilitas_tdk_susut;

  return nilai_jpb;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
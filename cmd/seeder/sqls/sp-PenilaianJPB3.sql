CREATE OR REPLACE FUNCTION "public"."penilaian_jpb3"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_jpb numeric := 0;
  tinggi_kolom numeric := 0;
	lebar_bentang numeric := 0;
	nilai_komponen_utama numeric := 0;
	atap char(1);
	lantai char(1);
	langit_langit char(1);
	dinding char(1);
	jml_lantai_bng int := 1;
	nl_atap numeric := 0;
	nilai_atap numeric := 0;
	nilai_lantai numeric := 0;
	nilai_langit_langit numeric := 0;
	nilai_material numeric := 0;
	type_konstruksi char(1);
	cari_dinding char(2);
	nilai_dinding numeric := 0;
	nilai_daya_dukung numeric:= 0;
	nilai_fasilitas numeric:= 0;
	nilai_total_per_m2 numeric:= 0;
	luas_bng numeric := 0;
	nilai_total_kali_luas numeric := 0;
	keliling_dinding numeric := 0;
	total_nilai_dinding numeric := 0;
	luas_mezzanine numeric := 0;
	nilai_mezzanine numeric:= 0;
	nilai_total_mezzanine numeric := 0;
	nilai_fasilitas_susut numeric := 0;
	nilai_sebelum_susut numeric := 0;
	tahun_dibangun char(4);
	tahun_renovasi char(4);
	umur_efektif numeric := 0;
	biaya_pengganti_baru numeric := 0;
	kondisi_bng char(1);
	persentase_penyusutan numeric := 0;
	nilai_setelah_susut numeric := 0;
	nilai_fasilitas_tdk_susut numeric:= 0;

begin

  -- Menentukan biaya komponen utama / m2
  BEGIN
	  SELECT "TinggiKolom3", "LebarBentang3", "KelilingDinding3", "LuasMezzanine3", "TipeKonstruksi"
	  INTO tinggi_kolom, lebar_bentang, keliling_dinding, luas_mezzanine, type_konstruksi
	  FROM "Jpb3"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" = NoUrut AND
      "JenisOp" = JenisOp	AND
      "NoBangunan" = NoBangunan;
  
  END;

  BEGIN
	  SELECT "NilaiDbkbJpb3"
	  INTO 	 nilai_komponen_utama
	  FROM 	 "DbkbJpb3"
	  WHERE "Provinsi_Kode" = Provinsi_Kode
      AND "Daerah_Kode" = Daerah_Kode
      AND "TahunDbkbJpb3" = Tahun
	  	AND "LebarBentukMinDbkbJpb3" <= lebar_bentang
	  	AND "LebarBentukMaxDbkbJpb3" >= lebar_bentang
	  	AND "TinggiKolomMinDbkbJpb3" <= tinggi_kolom
	  	AND "TinggiKolomMaxDbkbJpb3" >= tinggi_kolom;

    if not found then
      nilai_komponen_utama = 0;
    end if;
  END;

  -- Menentukan Biaya Komponen Material /m2
  BEGIN
	  SELECT "JenisAtap", "KodeLantai", "KodeLangitLangit", "LuasBangunan", "TahunDibangun", "TahunRenovasi",
      "Kondisi", "KodeDinding", "JmlLantaiBng"
	  INTO atap, lantai, langit_langit, luas_bng, tahun_dibangun, tahun_renovasi, kondisi_bng, dinding, jml_lantai_bng
	  FROM   "ObjekPajakBangunan"
	  WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" = NoUrut AND
      "JenisOp" = JenisOp	AND
      "NoBangunan" = NoBangunan AND
      "Jpb_Kode" = '03';
    
  END;

  -- Mencari Nilai Atap
  BEGIN
    SELECT "NilaiDbkbMaterial"
    INTO 	 nl_atap
    FROM 	 "DbkbMaterial"
    WHERE "Provinsi_Kode" = Provinsi_Kode
      AND "Daerah_Kode" = Daerah_Kode
      AND  "Thn_Dbkb_Material" = Tahun
      AND  "Pekerjaan_Kode" = '23'
      AND  "Kegiatan_Kode" = '0' || atap;

    if not found then
      nl_atap = 0;
    end if;
      
  END;
    
  nilai_atap := nl_atap / jml_lantai_bng;

  -- Mencari Nilai Lantai
  BEGIN
    SELECT "NilaiDbkbMaterial"
    INTO 	 nilai_lantai
    FROM 	 "DbkbMaterial"
    WHERE "Provinsi_Kode" = Provinsi_Kode
      AND "Daerah_Kode" = Daerah_Kode
      AND  "Thn_Dbkb_Material" = Tahun
      AND  "Pekerjaan_Kode" = '22'
      AND  "Kegiatan_Kode" = '0' || lantai;

    if not found then
      nilai_lantai = 0;
    end if;
      
  END;

  -- Mencari Nilai Langit-langit
  BEGIN
    SELECT "NilaiDbkbMaterial"
    INTO 	 nilai_langit_langit
    FROM 	 "DbkbMaterial"
    WHERE "Provinsi_Kode" = Provinsi_Kode
      AND "Daerah_Kode" = Daerah_Kode
      AND  "Thn_Dbkb_Material" = Tahun
      AND  "Pekerjaan_Kode" = '24'
      AND  "Kegiatan_Kode" = '0' || lantai;

    if not found then
      nilai_langit_langit = 0;
    end if;
      
  END;

  nilai_material := (nilai_atap + nilai_lantai + nilai_langit_langit) * (1.3);

  -- Menentukan Daya Dukung Lantai
  BEGIN
    SELECT "NilaiDbkbDayaDukung"
    INTO nilai_daya_dukung
    FROM "DbkbDayaDukung"
    WHERE "Provinsi_Kode" = Provinsi_Kode
      AND "Daerah_Kode" = Daerah_Kode
      AND "ThnDbkbDayaDukung" = tahun
      AND "TipeKonstruksi" = type_konstruksi;

    if not found then
      nilai_daya_dukung = 0;
    end if;

  END;

  -- Menentukan Fasilitas / m2
  BEGIN
    SELECT "fasilitas_susut_x_luas"
    INTO nilai_fasilitas
    FROM Fasilitas_Susut_X_Luas(
      Provinsi_Kode,
      Daerah_Kode,
      Kecamatan_Kode,
      Kelurahan_Kode,
      Blok_Kode,
      NoUrut,
      JenisOp,
      NoBangunan,
      '03',
      null,
      Tahun
    );

  END;

  -- Nilai Total / m2
  nilai_total_per_m2 := nilai_komponen_utama + nilai_material + nilai_daya_dukung + nilai_fasilitas;

  -- Menghitung Nilai Total / m2 dikali luas
  nilai_total_kali_luas := nilai_total_per_m2 * luas_bng;

  --  Menghitung Nilai dinding
  --  Jenis dinding	   		  		    Jenis Kegiatan Dinding
  --    1  >>  Kaca/Aluminium       01  >>  Kaca
  --    2  >>  Beton						    02	>>	Aluminium / Spandek
  --	  3  >>  Batu Bata/Conblok	  03	>>	Beton
  --	  4  >>  Kayu 						    04  >>  Batu Bata
  --	  5  >>  Seng						      05	>>	Kayu
  --	  6  >>  Tidak ada					  06  >>  Seng

  IF dinding = '1' THEN cari_dinding := '09'; -- bukan '01';
  ELSIF dinding = '2' THEN cari_dinding := '02';
  ELSIF dinding = '3' THEN cari_dinding := '03';
  ELSIF dinding = '4' THEN cari_dinding := '07';
  ELSIF dinding = '5' THEN cari_dinding := '08';
  ELSIF dinding = '6' THEN cari_dinding := null;
  END IF;

  -- Mencari Nilai dinding
  BEGIN
    SELECT "NilaiDbkbMaterial"
    INTO 	 nilai_dinding
    FROM 	 "DbkbMaterial"
    WHERE "Provinsi_Kode" = Provinsi_Kode
      AND "Daerah_Kode" = Daerah_Kode
      AND  "Thn_Dbkb_Material" = Tahun
      AND  "Pekerjaan_Kode" = '21'
      AND  "Kegiatan_Kode" = cari_dinding;

    if not found then
      nilai_dinding = 0;
    end if;

  END;
  total_nilai_dinding := (keliling_dinding * tinggi_kolom * (10/6) * nilai_dinding) * (1.3);

  -- Menghitung Nilai Mezzanine
  BEGIN
    SELECT "NilaiDbkbMezanin"
    INTO   nilai_mezzanine
    FROM   "DbkbMezanin"
    WHERE  "Provinsi_Kode" = Provinsi_Kode
      AND  "Daerah_Kode" = Daerah_Kode
      AND  "TahunDbkbMezanin" = Tahun;

    if not found then
      nilai_mezzanine = 0;
    end if;

  END;
  nilai_total_mezzanine := luas_mezzanine * nilai_mezzanine;

  -- Menghitung Nilai Fasilitas yang disusutkan
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

  -- Menghitung Nilai sebelum disusutkan
  nilai_sebelum_susut := nilai_total_kali_luas + total_nilai_dinding + nilai_total_mezzanine + nilai_fasilitas_susut;

  -- Menghitung Penyusutan Bangunan
  BEGIN
    SELECT "susut"
    INTO persentase_penyusutan
    FROM Susut(
      CAST(Tahun AS int),
      CAST(tahun_dibangun AS int),
      CAST(tahun_renovasi AS int),
      kondisi_bng,
      nilai_sebelum_susut,
      luas_bng,
      0
    );

  END;

  -- Menghitung Nilai Setelah disusutkan
  persentase_penyusutan  := persentase_penyusutan / 100;
  nilai_setelah_susut    := nilai_sebelum_susut - (nilai_sebelum_susut * persentase_penyusutan);

  -- Menghitung Fasilitas Lain yang tidak disusutkan
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

  nilai_jpb := nilai_setelah_susut + nilai_fasilitas_tdk_susut;
    
  return nilai_jpb;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
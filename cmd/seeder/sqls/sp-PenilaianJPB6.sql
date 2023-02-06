CREATE OR REPLACE FUNCTION "public"."penilaian_jpb6"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_jpb numeric := 0;
  kls_jpb6 char(1);
	nilai_komponen_utama numeric := 0;
	thn_dibangun char(4);
  thn_renovasi char(4);
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
	nilai_fasilitas_susut numeric := 0;
	nilai_sebelum_susut numeric := 0;
	persentase_penyusutan numeric := 0;
	persen_susut numeric := 0;
	nilai_setelah_susut numeric := 0;
	nilai_fasilitas_tdk_susut numeric := 0;

begin

  -- Menentukan biaya komponen utama /m2
  BEGIN
	  SELECT "KelasBangunan6"
	  INTO kls_jpb6
	  FROM "Jpb6"
	  WHERE "Provinsi_Kode" = Provinsi_Kode
	    AND "Daerah_Kode" = Daerah_Kode
	  	AND "Kecamatan_Kode" = Kecamatan_Kode
	  	AND "Kelurahan_Kode" = Kelurahan_Kode
	  	AND "Blok_Kode" = Blok_Kode
	  	AND "NoUrut" = NoUrut
	  	AND "JenisOp" = JenisOp
	  	AND "NoBangunan" = NoBangunan;
    
  END;

  BEGIN
	  SELECT "NilaiDbkbJpb6"
	  INTO nilai_komponen_utama
	  FROM  "DbkbJpb6"
	  WHERE "Provinsi_Kode" = Provinsi_Kode
	    AND "Daerah_Kode" = Daerah_Kode
	  	AND "TahunDbkbJpb6" = Tahun
	  	AND "KelasDbkbJpb6" = kls_jpb6;

    if not found then
      nilai_komponen_utama = 0;
    end if;
  END;

  -- Mengambil luas bangunan, tahun, dsb.
  BEGIN
    SELECT "LuasBangunan", "TahunDibangun", "TahunRenovasi", "Kondisi", "JmlLantaiBng"
    INTO luas_bng, thn_dibangun, thn_renovasi, kondisi_bng, jml_lantai_bng
	  FROM "ObjekPajakBangunan"
    WHERE "Provinsi_Kode" = Provinsi_Kode
      AND "Daerah_Kode" = Daerah_Kode
      AND "Kecamatan_Kode" = Kecamatan_Kode
      AND "Kelurahan_Kode" = Kelurahan_Kode
      AND "Blok_Kode" = Blok_Kode
      AND "NoUrut" = NoUrut
      AND "JenisOp" = JenisOp
      AND "NoBangunan" = NoBangunan
      AND "Jpb_Kode" = '06';
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
      '06',
      kls_jpb6,
      Tahun
    );

  END;

  -- Nilai Total / m2
  nilai_total_per_m2 := nilai_komponen_utama + nilai_fasilitas;

  -- Menghitung Nilai Total / m2 dikali luas
  nilai_total_kali_luas := nilai_total_per_m2 * luas_bng;

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
  nilai_sebelum_susut := nilai_total_kali_luas + nilai_fasilitas_susut;

  -- Menghitung Penyusutan Bangunan
  BEGIN
    SELECT "susut"
    INTO persentase_penyusutan
    FROM Susut(
      CAST(Tahun AS int),
      CAST(thn_dibangun AS int),
      CAST(thn_renovasi AS int),
      kondisi_bng,
      nilai_sebelum_susut,
      luas_bng,
      0
    );

  END;

	persen_susut := round((persentase_penyusutan / 100), 2);

  -- Menghitung Nilai Setelah disusutkan
  nilai_setelah_susut := nilai_sebelum_susut - (nilai_sebelum_susut * persen_susut);

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
  
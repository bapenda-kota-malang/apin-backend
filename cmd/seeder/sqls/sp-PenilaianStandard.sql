CREATE OR REPLACE FUNCTION "public"."penilaian_standard"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_standard numeric := 0;
  kd_jpb char(2);
  thn_dibangun char(4);
  thn_renovasi char(4);
  luas_bng numeric := 0;
  kondisi_bng char(1);
  jns_konstruksi_bng char(1);
  jml_lantai_bng int := 1;
  komp_utama 		numeric := 0;
  komp_material 		numeric := 0;
  fasilitas_x_luas 		numeric := 0;
  fasilitas_susut 		numeric := 0;
  fasilitas_tdk_susut 		numeric := 0;
  nilai1     		numeric := 0;
  nilai_sbl_susut  numeric := 0;
  nilai_stl_susut  numeric := 0;
  besar_susut      numeric := 0;
  persen_susut     numeric := 0;

begin

  SELECT "Jpb_Kode", "LuasBangunan", "TahunDibangun", "TahunRenovasi", "Kondisi", "JmlLantaiBng", "JenisKonstruksi"
  INTO kd_jpb, luas_bng, thn_dibangun, thn_renovasi, kondisi_bng, jml_lantai_bng, jns_konstruksi_bng
  FROM "ObjekPajakBangunan"
  WHERE "Provinsi_Kode" = Provinsi_Kode AND
    "Daerah_Kode" = Daerah_Kode AND
    "Kecamatan_Kode" = Kecamatan_Kode AND
    "Kelurahan_Kode" = Kelurahan_Kode AND
    "Blok_Kode" = Blok_Kode AND
    "NoUrut" = NoUrut AND
    "JenisOp" = JenisOp	AND
    "NoBangunan" = NoBangunan;

  if found then
    -- jika bangunan lain-lain, maka masukkan bangunan perumahan
	   IF kd_jpb IN ('01','10','11') THEN 
      kd_jpb := '01';
    ELSIF kd_jpb IN ('02','04','07','09') THEN
      kd_jpb := '02';
    END IF;
    
    IF thn_renovasi IS NULL OR thn_renovasi = ' ' THEN
      thn_renovasi := '0';
    END IF;

    -- cari biaya komponen utama / m2
    BEGIN
      SELECT "komp_utama_std"
      INTO komp_utama
      FROM Komp_Utama_STD(
        Provinsi_Kode,
        Daerah_Kode,
        Kecamatan_Kode,
        Kelurahan_Kode,
        Blok_Kode,
        NoUrut,
        JenisOp,
        NoBangunan,
        Tahun,
        kd_jpb,
        luas_bng,
        jml_lantai_bng,
        jns_konstruksi_bng
      );

    END;

    -- cari biaya komponen material / m2
    BEGIN
      SELECT "komp_material"
      INTO komp_material
      FROM Komp_Material(
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

    -- dirubah/ditambah TEGUH atas periksa PAK RUSLAN
    -- Tgl. 19/10/2000
    -- cari biaya fasilitas yang disusutkan ( yang dipengaruhi luas bangunan)
    BEGIN
      SELECT "fasilitas_susut_x_luas"
      INTO fasilitas_x_luas
      FROM Fasilitas_Susut_X_Luas(
        Provinsi_Kode,
        Daerah_Kode,
        Kecamatan_Kode,
        Kelurahan_Kode,
        Blok_Kode,
        NoUrut,
        JenisOp,
        NoBangunan,
        Jpb_Kode,
        '0',
        Tahun
      );

    END;

    nilai1 := (komp_utama + komp_material + fasilitas_x_luas) * Nluas_bng;

    BEGIN
      SELECT "fasilitas_susut"
      INTO fasilitas_susut
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

    nilai_sbl_susut := nilai1 + fasilitas_susut;
    BEGIN
      SELECT "susut"
      INTO besar_susut
      FROM Susut(
        CAST(Tahun AS int),
        CAST(thn_dibangun AS int),
        CAST(thn_renovasi AS int),
        kondisi_bng,
        nilai_sbl_susut,
        luas_bng,
        1
      );

    END;

    IF besar_susut IS NOT NULL OR besar_susut = 0 THEN
      persen_susut 	  := round((besar_susut / 100),2);
      nilai_stl_susut := (nilai_sbl_susut - (nilai_sbl_susut * persen_susut));
    ELSE
      nilai_stl_susut := nilai_sbl_susut;
    END IF;

    BEGIN
      SELECT "fasilitas_tidak_susut"
      INTO fasilitas_tdk_susut
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
    
    nilai_standard := nilai_stl_susut + fasilitas_tdk_susut;
  
  ELSE
    nilai_standard := 0;
  end if;

  return nilai_standard;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
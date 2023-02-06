CREATE OR REPLACE FUNCTION "public"."penilaian_bangunan"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "jpb_kode" bpchar, "luas_bangunan" numeric, "jmllantai" int4, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_bng numeric := 0;
begin
  -- DIBUAT OLEH : MADE
  -- TANGGAL     :
  -- REVISI KE   : 2
  -- DIREVISI    : TEGUH
  -- TGL. REVISI : 10-10-2000

  ----------------------------------------------------------------------------------------
  -- procedure ini digunakan untuk mencari nilai sistem bangunan masing-masing bangunan --
  --             yang dimiliki oleh op anggota/anak dari suatu NOP tertentu             --
  ----------------------------------------------------------------------------------------

  IF (Jpb_Kode IN ('02','04','05','07','09')) AND ((Luas_Bangunan > 1000) OR (JmlLantai > 4)) THEN

    ----------------------------------------------------------------------------------------
    -- jika kode jpb-nya = 02 (perkantoran swasta), maka panggil procedure PENILAIAN_JPB2 --
    ----------------------------------------------------------------------------------------
    IF Jpb_Kode = '02' THEN
      BEGIN
        SELECT "penilaian_jpb2"
        INTO nilai_bng
        FROM Penilaian_JPB2(
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

    ---------------------------------------------------------------------------------------
    -- jika kode jpb-nya = 04 (toko/apotik/pasar/ruko), panggil procedure PENILAIAN_JPB4 --
    ---------------------------------------------------------------------------------------
    ELSIF Jpb_Kode = '04' THEN
      BEGIN
        SELECT "penilaian_jpb4"
        INTO nilai_bng
        FROM Penilaian_JPB4(
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

    -----------------------------------------------------------------------------------
    -- jika kode jpb-nya = 05 (rumah sakit/klinik), panggil procedure PENILAIAN_JPB5 --
    -----------------------------------------------------------------------------------
    ELSIF Jpb_Kode = '05' THEN
      BEGIN
        SELECT "penilaian_jpb5"
        INTO nilai_bng
        FROM Penilaian_JPB5(
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

    -------------------------------------------------------------------------------------
    -- jika kode jpb-nya = 07 (hotel/restoran/wisma), panggil procedure PENILAIAN_JPB7 --
    -------------------------------------------------------------------------------------
    ELSIF Jpb_Kode = '07' THEN
      BEGIN
        SELECT "penilaian_jpb7"
        INTO nilai_bng
        FROM Penilaian_JPB7(
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

    -----------------------------------------------------------------------------------------
    -- jika kode jpb-nya = 09 (gedung pemerintahan), maka panggil procedure PENILAIAN_JPB9 --
    -----------------------------------------------------------------------------------------
    ELSIF Jpb_Kode = '09' THEN
      BEGIN
        SELECT "penilaian_jpb9"
        INTO nilai_bng
        FROM Penilaian_JPB9(
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
    END IF;

  ---------------------------------------------------------------------------------------
  -- jika kode jpb-nya = 03 (pabrik), panggil procedure op non standard PENILAIAN_JPB3 --
  ---------------------------------------------------------------------------------------
  ELSIF Jpb_Kode = '03' THEN
    BEGIN
      SELECT "penilaian_jpb3"
      INTO nilai_bng
      FROM Penilaian_JPB3(
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

  --------------------------------------------------------------------------------------------------
  -- jika kode jpb-nya = 06 (olahraga/rekreasi), panggil procedure op non standard PENILAIAN_JPB6 --
  --------------------------------------------------------------------------------------------------
  ELSIF Jpb_Kode = '06' THEN
    BEGIN
      SELECT "penilaian_jpb6"
      INTO nilai_bng
      FROM Penilaian_JPB6(
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

  ---------------------------------------------------------------------------------------------------------
  -- jika kode jpb-nya = 08 (bengkel/gudang/pertanian), panggil procedure op non standard PENILAIAN_JPB8 --
  ---------------------------------------------------------------------------------------------------------
  ELSIF Jpb_Kode = '08' THEN
    BEGIN
      SELECT "penilaian_jpb8"
      INTO nilai_bng
      FROM Penilaian_JPB8(
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

  ----------------------------------------------------------------------------
  -- jika kode jpb-nya = 11 (bangunan tidak kena pajak), nilai bangunan = 0 --
  ----------------------------------------------------------------------------
  ELSIF Jpb_Kode = '11' THEN
    nilai_bng := 0;
  
  -------------------------------------------------------------------------------------------------
  -- jika kode jpb-nya = 12 (bangunan parkir), panggil procedure op non standard PENILAIAN_JPB12 --
  -------------------------------------------------------------------------------------------------
  ELSIF Jpb_Kode = '12' THEN
    BEGIN
      SELECT "penilaian_jpb12"
      INTO nilai_bng
      FROM Penilaian_JPB12(
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
  
  ------------------------------------------------------------------------------------------
  -- jika kode jpb-nya = 13 (apartemen), panggil procedure op non standard PENILAIAN_JPB13 --
  ------------------------------------------------------------------------------------------
  ELSIF Jpb_Kode = '13' THEN
    BEGIN
      SELECT "penilaian_jpb13"
      INTO nilai_bng
      FROM Penilaian_JPB13(
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

  ----------------------------------------------------------------------------------------------
  -- jika kode jpb-nya = 14 (pompa bensin), panggil procedure op non standard PENILAIAN_JPB14 --
  ----------------------------------------------------------------------------------------------
  ELSIF Jpb_Kode = '14' THEN
    BEGIN
      SELECT "penilaian_jpb14"
      INTO nilai_bng
      FROM Penilaian_JPB14(
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

  -----------------------------------------------------------------------------------------------
  -- jika kode jpb-nya = 15 (tangki minyak), panggil procedure op non standard PENILAIAN_JPB15 --
  -----------------------------------------------------------------------------------------------
  ELSIF Jpb_Kode = '15' THEN
    BEGIN
      SELECT "penilaian_jpb15"
      INTO nilai_bng
      FROM Penilaian_JPB15(
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

  -----------------------------------------------------------------------------------------------
  -- jika kode jpb-nya = 16(gedung sekolah), panggil procedure op non standard PENILAIAN_JPB16 --
  -----------------------------------------------------------------------------------------------
  ELSIF Jpb_Kode = '16' THEN
    BEGIN
      SELECT "penilaian_jpb16"
      INTO nilai_bng
      FROM Penilaian_JPB16(
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

  ------------------------------------------------------------------------------------------
  -- jika bukan semuanya, maka panggil procedure penilaian op standard PENILAIAN_STANDARD --
  ------------------------------------------------------------------------------------------
  ELSE
    BEGIN
      SELECT "penilaian_standard"
      INTO nilai_bng
      FROM Penilaian_Standard(
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
  END IF;

  nilai_bng := ROUND(nilai_bng, 0);

  return nilai_bng;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
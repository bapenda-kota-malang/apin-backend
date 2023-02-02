DROP FUNCTION IF EXISTS "public"."penilaian_jpb5"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar);
CREATE OR REPLACE FUNCTION "public"."penilaian_jpb5"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_jpb numeric := 0;
  nilai_dbkb_jpb5 numeric := 0;
  penyusutan numeric := 0;
  thn_dibangun char(4);
  thn_renovasi char(4);
  luas_bng numeric := 0;
  kondisi_bng char(1);
  jml_lantai_bng int := 1;
  kls_jpb5 char(1);
  luas_kmr_jpb5 numeric := 0;
  luas_rng_lain numeric:= 0;
  jml_satuan numeric:= 0;
  status_fasilitas  varchar(4);
  kd_fasilitas char(2);
  ketergantungan char(2);
  nilai_fasilitas numeric := 0;
  nilai_satuan numeric := 0;
  nilai_sebelum_susut numeric := 0;
  rec_fasilitas record;
  cur_fasilitas cursor for SELECT "Status", "Fasilitas_Kode", "Ketergantungan" FROM "Fasilitas" WHERE "Status" IN ('0','2','3');

begin

  BEGIN
    SELECT a."TahunDibangun", a."TahunRenovasi", a."LuasBangunan", a."JmlLantaiBng",
      a."Kondisi", b."KelasBangunan5", b."LuasKamarAcCentral", b."LuasRuangLainAcCentral"
    INTO thn_dibangun, thn_renovasi, luas_bng, jml_lantai_bng, kondisi_bng,
      kls_jpb5, luas_kmr_jpb5, luas_rng_lain
    FROM "ObjekPajakBangunan" AS a, "Jpb5" AS b
    WHERE a."Provinsi_Kode" = Provinsi_Kode
      AND a."Daerah_Kode" = Daerah_Kode
      AND a."Kecamatan_Kode" = Kecamatan_Kode
      AND a."Kelurahan_Kode" = Kelurahan_Kode
      AND a."Blok_Kode" = Blok_Kode
      AND a."NoUrut" = NoUrut
      AND a."JenisOp" = JenisOp
      AND a."NoBangunan" = NoBangunan
      AND a."Jpb_Kode" = '05'
      AND b."Provinsi_Kode" = a."Provinsi_Kode"
      AND b."Daerah_Kode" = a."Daerah_Kode"
      AND b."Kecamatan_Kode" = a."Kecamatan_Kode"
      AND b."Kelurahan_Kode" = a."Kelurahan_Kode"
      AND b."Blok_Kode" = a."Blok_Kode"
      AND b."NoUrut" = a."NoUrut"
      AND b."JenisOp" = a."JenisOp"
      AND b."NoBangunan" = a."NoBangunan";
      
  END;

  -- cari nilai komponen utama / m2
  BEGIN
    SELECT "NilaiDbkbJpb5"
    INTO nilai_dbkb_jpb5
    FROM "DbkbJpb5"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "TahunDbkbJpb5" = Tahun AND
      "KelasDbkbJpb5" = kls_jpb5 AND
      "LantaiMinJpb5" <= jml_lantai_bng AND
      "LantaiMaxJpb5" >= jml_lantai_bng;
    
  END;

  -- cari nilai komponen utama X luas bangunan
  nilai_dbkb_jpb5 := nilai_dbkb_jpb5 * luas_bng;

  -- cari nilai fasilitas yang dipengaruhi luas, jumlah kamar, luas kamar
  open cur_fasilitas;

    loop
    fetch cur_fasilitas into rec_fasilitas;
    -- exit when no more row to fetch
    exit when not found;

    IF rec_fasilitas."Status" = '0' THEN
      BEGIN
        SELECT "JumlahSatuan"
        INTO jml_satuan
        FROM "FasilitasBangunan"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Kecamatan_Kode" = Kecamatan_Kode AND
          "Kelurahan_Kode" = Kelurahan_Kode AND
          "Blok_Kode" = Blok_Kode AND
          "NoUrut" = NoUrut AND
          "JenisOp" = JenisOp AND
          "NoBangunan" = NoBangunan AND
          "KodeFasilitas" = kd_fasilitas;

        if not found then
          jml_satuan = 0;
        end if;
      END;
	  END IF;

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
          "Jpb_Kode" = '05' AND
          "KlsBintang" = kls_jpb5;

        if not found then
          nilai_satuan = 0;
        end if;
      END;
	  ELSE
      nilai_satuan := 0;
	  END IF;

    -- penentuan nilai fasilitas yang dipengaruhi luas, jumlah kamar, luas kamar
    IF rec_fasilitas."Status" = '0'  THEN
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * jml_satuan * luas_bng);
    ELSIF rec_fasilitas."Status" = '2'  THEN
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * luas_kmr_jpb5);
    ELSIF rec_fasilitas."Status" = '3'  THEN
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * luas_rng_lain);
    ELSE 
      nilai_fasilitas := nilai_fasilitas;
    END IF;
    
  end loop;
  
  close cur_fasilitas;

  -- cari nilai komponen utama + nilai fasilitas
  nilai_jpb := nilai_dbkb_jpb5 + nilai_fasilitas;

  -- cari nilai fasilitas yang disusutkan
  BEGIN
    SELECT "fasilitas_susut"
    INTO nilai_fasilitas
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
  nilai_sebelum_susut := nilai_jpb + nilai_fasilitas;

  -- cari prosentase penyusutan
  BEGIN
    SELECT "susut"
    INTO penyusutan
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

  -- cari nilai setelah disusutkan
  IF (penyusutan IS NOT NULL) OR (penyusutan = 0) THEN
    penyusutan := round((penyusutan / 100),2);
    nilai_jpb := nilai_jpb - (nilai_jpb * penyusutan);
  ELSE
	  nilai_jpb := nilai_jpb;
  END IF;

  -- cari nilai fasilitas yang tidak disusutkan
  BEGIN
    SELECT "fasilitas_tidak_susut"
    INTO nilai_fasilitas
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

  -- cari nilai jpb5
  nilai_jpb := nilai_jpb + nilai_fasilitas;

  return nilai_jpb;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
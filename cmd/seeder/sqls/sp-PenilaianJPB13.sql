CREATE OR REPLACE FUNCTION "public"."penilaian_jpb13"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_jpb numeric := 0;
  nilai_jpb13 numeric := 0;
  penyusutan numeric := 0;
  thn_dibangun_bng char(4);
  thn_renovasi_bng char(4);
  luas_bng numeric := 0;
  jml_lantai_bng int := 1;
  kondisi_bng char(1);
  nilai_dbkb_jpb13 numeric := 0;
  kls_jpb13 char(1);
  jml_jpb13 numeric := 0;
  luas_jpb13 bigint;
  luas_jpb13_lain bigint;
  jml_satuan numeric:= 0;
  kd_fasilitas char(2);
  status_fasilitas  varchar(4);
  ketergantungan char(2);
  nilai_fasilitas  numeric := 0;
  nilai_satuan     numeric := 0;
  nilai_sebelum_susut numeric := 0;
  rec_fasilitas record;
  cur_fasilitas cursor for SELECT "Status", "Fasilitas_Kode", "Ketergantungan" FROM "Fasilitas" WHERE "Status" IN ('0','1','2','3');

begin

  -- cari data bangunan
  BEGIN
    SELECT a."TahunDibangun", a."TahunRenovasi", a."LuasBangunan", a."Kondisi", a."JmlLantaiBng",
      b."KelasBangunan13", b."JumlahApartment", b."LuasApartAcCentral", b."LuasRuangLainAcCentral"
    INTO thn_dibangun_bng, thn_renovasi_bng, luas_bng, kondisi_bng, jml_lantai_bng,
      kls_jpb13, jml_jpb13, luas_jpb13, luas_jpb13_lain
    FROM "ObjekPajakBangunan" AS a, "Jpb13" AS b
    WHERE a."Provinsi_Kode" = Provinsi_Kode
      AND a."Daerah_Kode" = Daerah_Kode
      AND a."Kecamatan_Kode" = Kecamatan_Kode
      AND a."Kelurahan_Kode" = Kelurahan_Kode
      AND a."Blok_Kode" = Blok_Kode
      AND a."NoUrut" = NoUrut
      AND a."JenisOp" = JenisOp
      AND a."NoBangunan" = NoBangunan
      AND a."Jpb_Kode" = '13'
      AND b."Provinsi_Kode" = a."Provinsi_Kode"
      AND b."Daerah_Kode" = a."Daerah_Kode"
      AND b."Kecamatan_Kode" = a."Kecamatan_Kode"
      AND b."Kelurahan_Kode" = a."Kelurahan_Kode"
      AND b."Blok_Kode" = a."Blok_Kode"
      AND b."NoUrut" = a."NoUrut"
      AND b."JenisOp" = a."JenisOp"
      AND b."NoBangunan" = a."NoBangunan";

  END;

  -- cari nilai komponen utama dbkb jpb 13 dari tabel DBKB_JPB13
  BEGIN
    SELECT "NilaiDbkbJpb13"
    INTO nilai_dbkb_jpb13
    FROM "DbkbJpb13"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
    "Daerah_Kode" = Daerah_Kode AND
    "TahunDbkbJpb13" = Tahun AND
    "KelasDbkbJpb13" = kls_jpb13 AND
    "LantaiMinJpb13" <= jml_lantai_bng AND
    "LantaiMaxJpb13" >= jml_lantai_bng;
    
    if not found then
      nilai_dbkb_jpb13 := 0;
    end if;
  END;

  -- nilai komponen utama X luas bangunan
  nilai_dbkb_jpb13 := nilai_dbkb_jpb13 * luas_bng;

  -- cari nilai fasilitas yang dipengaruhi luas, jumlah unit apartemen
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

    IF rec_biaya_fasilitas."Ketergantungan" = '0' THEN
      BEGIN
        SELECT "nilai_non_dep" INTO nilai_satuan
        FROM "fas_non_dep"
        WHERE "kd_propinsi" = Provinsi_Kode AND
          "kd_dati2" = Daerah_Kode AND
          "thn_non_dep" = Tahun AND
          "KodeFasilitas" = rec_biaya_fasilitas."Fasilitas_Kode";

        if not found then
          nilai_satuan = 0;
        end if;
      END;
	  ELSIF rec_biaya_fasilitas."Ketergantungan" = '1' THEN
      BEGIN
        SELECT "nilai_dep_min_max" INTO nilai_satuan
            FROM "fas_dep_min_max"
            WHERE "kd_propinsi" = Provinsi_Kode AND
              "kd_dati2" = Daerah_Kode AND
              "thn_dep_min_max" = Tahun AND
              "KodeFasilitas" = rec_biaya_fasilitas."Fasilitas_Kode" AND
              "kls_dep_min" <= JmlLantai AND
              "kls_dep_max" >= JmlLantai;
            
            if not found then
              nilai_satuan = 0;
            end if;
      END;
    ELSIF rec_biaya_fasilitas."Ketergantungan" = '2' THEN
      BEGIN
        SELECT "nilai_fasilitas_kls_bintang" INTO nilai_satuan
            FROM "fas_dep_jpb_kls_bintang"
            WHERE "kd_propinsi" = Provinsi_Kode AND
              "kd_dati2" = Daerah_Kode AND
              "thn_dep_jpb_kls_bintang" = Tahun AND
              "KodeFasilitas" = rec_biaya_fasilitas."Fasilitas_Kode" AND
              "kd_jpb" = '13' AND
              "kls_bintang" = kls_jpb13;
            
            if not found then
              nilai_satuan = 0;
            end if;
      END;
	  ELSE
      nilai_satuan := 0;
	  END IF;

    IF status_fasilitas = '0' THEN
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * jml_satuan * luas_bng);
    ELSIF status_fasilitas = '1' THEN
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * jml_jpb13);
    ELSIF status_fasilitas = '2' THEN
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * luas_jpb13);
    ELSIF status_fasilitas = '3' THEN
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * luas_jpb13_lain);
    ELSE
      nilai_fasilitas := nilai_fasilitas;
    END IF;
    
  end loop;
  
  close cur_fasilitas;

  -- nilai komponen utama + nilai fasilitas
  nilai_jpb13 := nilai_dbkb_jpb13 + nilai_fasilitas;

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
  
  nilai_jpb13     := nilai_jpb13 + nilai_fasilitas;

  -- cari prosentase penyusutan
  BEGIN
    SELECT "susut"
    INTO penyusutan
    FROM Susut(
      CAST(Tahun AS int),
      CAST(thn_dibangun_bng AS int),
      CAST(thn_renovasi_bng AS int),
      kondisi_bng,
      nilai_jpb13,
      luas_bng,
      0
    );

  END;

  -- cari nilai setelah disusutkan
  penyusutan := ROUND((penyusutan * nilai_jpb13) / 100);
  nilai_jpb13 := nilai_jpb13 - penyusutan;

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

  nilai_jpb13 := nilai_jpb13 + nilai_fasilitas;
  nilai_jpb := nilai_jpb13;

  return nilai_jpb;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
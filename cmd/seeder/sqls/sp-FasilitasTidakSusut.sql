DROP FUNCTION IF EXISTS "public"."fasilitas_tidak_susut"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar);
CREATE OR REPLACE FUNCTION "public"."fasilitas_tidak_susut"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_fasilitas numeric := 0;
  kd_fasilitas char(2);
  ketergantungan char(2);
  jml_satuan bigint := 0;
  nilai_satuan numeric := 0;
  rec_fas_tdk_susut record;
  cur_fas_tdk_susut cursor for SELECT "Fasilitas_Kode", "Ketergantungan"
    FROM "Fasilitas"
    WHERE "Status" = '5';

begin

  -- proses c_fas1
  open cur_fas_tdk_susut;

  loop
    fetch cur_fas_tdk_susut into rec_fas_tdk_susut;
    -- exit when no more row to fetch
    exit when not found;

    -- cari nilai satuan untuk masing2 kode fasilitas dari tabel FAS_NON_DEP
    BEGIN
      SELECT "NilaiNonDep" INTO nilai_satuan
      FROM "FasNonDep"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
        "Daerah_Kode" = Daerah_Kode AND
        "Thn_Non_Dep" = Tahun AND
        "Fasilitas_Kode" = rec_fas_tdk_susut."Fasilitas_Kode";
      
      if not found then
        nilai_satuan := 0;
      end if;
    END;

    -- cari jumlah satuan untuk masing2 kode fasilitas dari tabel DAT_FASILITAS_BANGUNAN
    BEGIN
      SELECT "JumlahSatuan" INTO jml_satuan
      FROM "FasilitasBangunan"
      WHERE "Provinsi_Kode" = Provinsi_Kode AND
        "Daerah_Kode" =  Daerah_Kode AND
        "Kecamatan_Kode" = Kecamatan_Kode AND
        "Kelurahan_Kode" = Kelurahan_Kode AND
        "Blok_Kode" = Blok_Kode AND
        "NoUrut" = NoUrut AND
        "JenisOp" = JenisOp AND
        "NoBangunan" = NoBangunan AND
        "KodeFasilitas" = rec_fas_tdk_susut."Fasilitas_Kode";

      if not found then
        jml_satuan := 0;
      end if;
    END;

    -- jika kode fasilitas = '44' [listrik] maka biaya fasilitas listrik / 1000
    IF kd_fasilitas = '44' THEN
      nilai_fasilitas := nilai_fasilitas + ((nilai_satuan * jml_satuan)/1000);
    ELSE
      nilai_fasilitas := nilai_fasilitas + (nilai_satuan * jml_satuan);
    END IF;
  end loop;
  
  close cur_fas_tdk_susut;

  return nilai_fasilitas;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
CREATE OR REPLACE FUNCTION "public"."fasilitas_susut"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "jmllantai" int4, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_fasilitas numeric := 0;
  jml_satuan bigint := 0;
  kd_fasilitas char(2);
  ketergantungan char(2);
  nilai_satuan numeric := 0;
  rec_biaya_fasilitas record;
  cursor_biaya_fasilitas cursor for SELECT "Fasilitas_Kode", "Ketergantungan" FROM "Fasilitas" WHERE "Status" = '4';
begin

  -- proses cursor_biaya_fasilitas
  open cursor_biaya_fasilitas;

  loop
    fetch cursor_biaya_fasilitas into rec_biaya_fasilitas;
    -- exit when no more row to fetch
    exit when not found;

    -- cari jumlah satuan untuk masing2 kode fasilitas dari tabel "FasilitasBangunan"
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

    -- cari nilai satuan untuk masing2 kode fasilitas --
    IF ketergantungan = '0' THEN
      BEGIN
        SELECT "NilaiNonDep" INTO nilai_satuan
        FROM "FasNonDep"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Thn_Non_Dep" = Tahun AND
          "Fasilitas_Kode" = rec_biaya_fasilitas."Fasilitas_Kode";

        if not found then
          nilai_satuan = 0;
        end if;
      END;
	  ELSIF ketergantungan = '1' THEN
      BEGIN
        IF rec_biaya_fasilitas."Fasilitas_Kode" IN ('30','31','32') THEN
          BEGIN
            SELECT "NilaiDepMinMax" INTO nilai_satuan
            FROM "FasDepMinMax"
            WHERE "Provinsi_Kode" = Provinsi_Kode AND
              "Daerah_Kode" = Daerah_Kode AND
              "TahunDepMinMax" = Tahun AND
              "Fasilitas_Kode" = rec_biaya_fasilitas.kd_fasilitas AND
              "KlsDepMin" <= JmlLantai AND
              "KlsDepMax" >= JmlLantai;
            
            if not found then
              nilai_satuan = 0;
            end if;
          END;
        ELSE
          BEGIN
            SELECT "NilaiDepMinMax" INTO nilai_satuan
            FROM "FasDepMinMax"
            WHERE "Provinsi_Kode" = Provinsi_Kode AND
              "Daerah_Kode" = Daerah_Kode AND
              "TahunDepMinMax" = Tahun AND
              "Fasilitas_Kode" = rec_biaya_fasilitas."Fasilitas_Kode" AND
              "KlsDepMin" <= jml_satuan AND
              "KlsDepMax" >= jml_satuan;
          
            if not found then
              nilai_satuan = 0;
            end if;
          END;
        END IF;
      END;
	  ELSE
      nilai_satuan := 0;
	  END IF;

	  nilai_fasilitas := nilai_fasilitas + (nilai_satuan * jml_satuan);

  end loop;
  
  close cursor_biaya_fasilitas;
  
  return nilai_fasilitas;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
DROP FUNCTION IF EXISTS "public"."fasilitas_susut_x_luas"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "kd_jpb" bpchar, "kls_bintang" bpchar, "tahun" varchar);
CREATE OR REPLACE FUNCTION "public"."fasilitas_susut_x_luas"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobangunan" int8, "kd_jpb" bpchar, "kls_bintang" bpchar, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_fasilitas numeric := 0;
  kd_fasilitas char(2);
  jml_satuan bigint := 0;
  nilai numeric := 0;
  nilai_fas numeric := 0;
  rec_fas1 record;
  c_fas1 cursor for SELECT "KodeFasilitas", "JumlahSatuan", "Status", "Ketergantungan"
    FROM "FasilitasBangunan", "Fasilitas"
    WHERE "Provinsi_Kode" = Provinsi_Kode AND
      "Daerah_Kode" = Daerah_Kode AND
      "Kecamatan_Kode" = Kecamatan_Kode AND
      "Kelurahan_Kode" = Kelurahan_Kode AND
      "Blok_Kode" = Blok_Kode AND
      "NoUrut" = NoUrut AND
      "JenisOp" = JenisOp AND
      "NoBangunan" = NoBangunan AND
      "Fasilitas_Kode" = "KodeFasilitas" AND
      "Status" = '0';

begin

  -- proses c_fas1
  open c_fas1;

  loop
    fetch c_fas1 into rec_fas1;
    -- exit when no more row to fetch
    exit when not found;
    
    IF rec_fas1."Ketergantungan" = '0' THEN
      BEGIN
        SELECT "NilaiNonDep"
        INTO nilai
        FROM "FasNonDep"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "Thn_Non_Dep" = Tahun AND
          "Fasilitas_Kode" = rec_fas1."KodeFasilitas";
          
        nilai_fas := nilai_fas + (nilai * rec_fas1."JumlahSatuan");
    	  if not found then
          nilai_fas := nilai_fas;
        end if;
      END;
    ELSIF rec_fas1."Ketergantungan" = '1' THEN
      BEGIN
        SELECT "NilaiDepMinMax"
        INTO nilai
        FROM "FasDepMinMax"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "TahunDepMinMax" = Tahun AND
          "Fasilitas_Kode" = rec_fas1."KodeFasilitas" AND
          "KlsDepMin" <= rec_fas1."JumlahSatuan" AND
          "KlsDepMax" >= rec_fas1."JumlahSatuan";
          
        nilai_fas := nilai_fas + (nilai * rec_fas1."JumlahSatuan");
        if not found then
          nilai_fas := nilai_fas;
        end if;
      END;
    ELSIF rec_fas1."Ketergantungan" = '2' THEN
      BEGIN
        SELECT "NilaiFasilitasKlsBintang"
        INTO nilai
        FROM "FasDepJpbKlsBintang"
        WHERE "Provinsi_Kode" = Provinsi_Kode AND
          "Daerah_Kode" = Daerah_Kode AND
          "TahunDep" = Tahun AND
          "Fasilitas_Kode" = rec_fas1."KodeFasilitas" AND
          "Jpb_Kode" = Kd_Jpb AND
          "KlsBintang" = Kls_Bintang;
          
        nilai_fas := nilai_fas + (nilai * rec_fas1."JumlahSatuan");
        
        if not found then
          nilai_fas := nilai_fas;
        end if;
      END;
    END IF;

    nilai_fasilitas := nilai_fas;
  end loop;
  
  close c_fas1;

  return nilai_fasilitas;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
  
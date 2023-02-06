CREATE OR REPLACE FUNCTION "public"."penilaian_bumi"("provinsi_kode" bpchar, "daerah_kode" bpchar, "kecamatan_kode" bpchar, "kelurahan_kode" bpchar, "blok_kode" bpchar, "nourut" bpchar, "jenisop" bpchar, "nobumi" int8, "kodeznt" bpchar, "luasbumi" int8, "tahun" varchar)
  RETURNS "pg_catalog"."numeric" AS $BODY$
declare
  nilai_bumi numeric := 0;
  nilai_nir numeric := 0;
begin

  if JenisOp IN ('7','8') THEN
    begin
      select "NilaiPerM2Tanah"
      into nilai_nir
      from "KelasTanah"
      where "KdTanah" = KodeZNT;

      if not found then
        nilai_nir = 0;
      end if;
		end;
	else
		begin
			select "Nir"
			into nilai_nir
			from "DataNir"
			where "Provinsi_Kode" = Provinsi_Kode
        and "Daerah_Kode" = Daerah_Kode
        and "Kecamatan_Kode" = Kecamatan_Kode
        and "Kelurahan_Kode" = Kelurahan_Kode
        and "Znt_Kode" = KodeZNT
        and "Tahun" = Tahun;
      
      if not found then
        nilai_nir = 0;
      end if;
		end;
	end if;

  -- cari nilai sistem bumi yaitu : (nir X 1000) dikalikan luas bumi
  nilai_bumi := LuasBumi * nilai_nir;

  return nilai_bumi;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
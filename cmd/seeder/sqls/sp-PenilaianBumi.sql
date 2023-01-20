-- Penilaian_Bumi
create or replace function Penilaian_Bumi(
  Provinsi_Kode char(2),
  Daerah_Kode	char(2),
  Kecamatan_Kode char(3),
  Kelurahan_Kode char(3),
  Blok_Kode char(3),
  NoUrut char(4),
  JenisOp char(1),
  NoBumi bigint,
  KodeZNT char(2),
  LuasBumi bigint,
  Tahun varchar(4)
)
returns numeric
language plpgsql    
as $$
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

end;$$
-- sample
-- SELECT * FROM Penilaian_Bumi('35','73','050','009','016','0351','1',1,'BZ',95,'2023')
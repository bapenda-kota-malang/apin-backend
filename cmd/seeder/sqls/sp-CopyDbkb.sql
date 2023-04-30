CREATE OR REPLACE FUNCTION "public"."copydbkb"("tahunbaru" varchar)
  RETURNS "pg_catalog"."varchar" AS $BODY$
declare
  tempTahun numeric := 0;
  tempTahun2 varchar(4);
begin
  tempTahun = CAST(TahunBaru AS int) - 1;
  tempTahun2 = CAST(tempTahun AS text);

  INSERT INTO "HargaSatuan" ("Propinsi_Id","Dati2_Id","TahunSatuan","Pekerjaan_Kode","Kegiatan_Kode","HargaSatuan","CreatedAt") 
  SELECT "HargaSatuan"."Propinsi_Id","HargaSatuan"."Dati2_Id",TahunBaru,"HargaSatuan"."Pekerjaan_Kode","HargaSatuan"."Kegiatan_Kode","HargaSatuan"."HargaSatuan",NOW() FROM "HargaSatuan" WHERE "HargaSatuan"."TahunSatuan" = tempTahun2;

  INSERT INTO "HargaResource" ("Propinsi_Id","Dati2_Id","TahunResource","GroupResource_Kode","Resource_Kode","Kanwil_Id","KPPBB_Id","JenisDokumen","NoDokumen","HargaResource","CreatedAt") 
  SELECT "HargaResource"."Propinsi_Id","HargaResource"."Dati2_Id",TahunBaru,"HargaResource"."GroupResource_Kode","HargaResource"."Resource_Kode","HargaResource"."Kanwil_Id","HargaResource"."KPPBB_Id","HargaResource"."JenisDokumen","HargaResource"."NoDokumen","HargaResource"."HargaResource",NOW() FROM "HargaResource" WHERE "HargaResource"."TahunResource" = tempTahun2;

  INSERT INTO "HargaKegiatan" ("Propinsi_Id","Dati2_Id","TahunKegiatan","JPB_Kode","Bangunan_Tipe","BangunanLantai_Kode","Pekerjaan_Kode","Kegiatan_Kode","HargaKegiatan","CreatedAt") 
  SELECT "HargaKegiatan"."Propinsi_Id","HargaKegiatan"."Dati2_Id",TahunBaru,"HargaKegiatan"."JPB_Kode","HargaKegiatan"."Bangunan_Tipe","HargaKegiatan"."BangunanLantai_Kode","HargaKegiatan"."Pekerjaan_Kode","HargaKegiatan"."Kegiatan_Kode","HargaKegiatan"."HargaKegiatan",NOW() FROM "HargaKegiatan" WHERE "HargaKegiatan"."TahunKegiatan" = tempTahun2;

  INSERT INTO "HargaKegiatanJPB8" ("Propinsi_Id","Dati2_Id","TahunKegiatanJPB8","Pekerjaan_Kode","Kegiatan_Kode","HRGLBRBENTMIN","HRGLBRBENTMAX","HRGTINGKOLOMMIN","HRGTINGKOLOMMAX","HargaKegiatanJPB8","CreatedAt") 
  SELECT "HargaKegiatanJPB8"."Propinsi_Id","HargaKegiatanJPB8"."Dati2_Id",TahunBaru,"HargaKegiatanJPB8"."Pekerjaan_Kode","HargaKegiatanJPB8"."Kegiatan_Kode","HargaKegiatanJPB8"."HRGLBRBENTMIN","HargaKegiatanJPB8"."HRGLBRBENTMAX","HargaKegiatanJPB8"."HRGTINGKOLOMMIN","HargaKegiatanJPB8"."HRGTINGKOLOMMAX","HargaKegiatanJPB8"."HargaKegiatanJPB8",NOW() FROM "HargaKegiatanJPB8" WHERE "HargaKegiatanJPB8"."TahunKegiatanJPB8" = tempTahun2;
  
  INSERT INTO "DbkbJpb2" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb2","KelasDbkbJpb2","LantaiMinJpb2","LantaiMaxJpb2","NilaiDbkbJpb2","CreatedAt") 
  SELECT "DbkbJpb2"."Provinsi_Kode","DbkbJpb2"."Daerah_Kode",TahunBaru,"DbkbJpb2"."KelasDbkbJpb2","DbkbJpb2"."LantaiMinJpb2","DbkbJpb2"."LantaiMaxJpb2","DbkbJpb2"."NilaiDbkbJpb2",NOW() FROM "DbkbJpb2" WHERE "DbkbJpb2"."TahunDbkbJpb2" = tempTahun2;
  
  INSERT INTO "DbkbJpb3" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb3","LebarBentukMinDbkbJpb3","LebarBentukMaxDbkbJpb3","TinggiKolomMinDbkbJpb3","TinggiKolomMaxDbkbJpb3","NilaiDbkbJpb3","CreatedAt") 
  SELECT "DbkbJpb3"."Provinsi_Kode","DbkbJpb3"."Daerah_Kode",TahunBaru,"DbkbJpb3"."LebarBentukMinDbkbJpb3","DbkbJpb3"."LebarBentukMaxDbkbJpb3","DbkbJpb3"."TinggiKolomMinDbkbJpb3","DbkbJpb3"."TinggiKolomMaxDbkbJpb3","DbkbJpb3"."NilaiDbkbJpb3",NOW() FROM "DbkbJpb3" WHERE "DbkbJpb3"."TahunDbkbJpb3" = tempTahun2;
  
  INSERT INTO "DbkbJpb4" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb4","KelasDbkbJpb4","LantaiMinJpb4","LantaiMaxJpb4","NilaiDbkbJpb4","CreatedAt") 
  SELECT "DbkbJpb4"."Provinsi_Kode","DbkbJpb4"."Daerah_Kode",TahunBaru,"DbkbJpb4"."KelasDbkbJpb4","DbkbJpb4"."LantaiMinJpb4","DbkbJpb4"."LantaiMaxJpb4","DbkbJpb4"."NilaiDbkbJpb4",NOW() FROM "DbkbJpb4" WHERE "DbkbJpb4"."TahunDbkbJpb4" = tempTahun2;
  
  INSERT INTO "DbkbJpb5" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb5","KelasDbkbJpb5","LantaiMinJpb5","LantaiMaxJpb5","NilaiDbkbJpb5","CreatedAt") 
  SELECT "DbkbJpb5"."Provinsi_Kode","DbkbJpb5"."Daerah_Kode",TahunBaru,"DbkbJpb5"."KelasDbkbJpb5","DbkbJpb5"."LantaiMinJpb5","DbkbJpb5"."LantaiMaxJpb5","DbkbJpb5"."NilaiDbkbJpb5",NOW() FROM "DbkbJpb5" WHERE "DbkbJpb5"."TahunDbkbJpb5" = tempTahun2;
  
  INSERT INTO "DbkbJpb6" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb6","KelasDbkbJpb6","NilaiDbkbJpb6","CreatedAt") 
  SELECT "DbkbJpb6"."Provinsi_Kode","DbkbJpb6"."Daerah_Kode",TahunBaru,"DbkbJpb6"."KelasDbkbJpb6","DbkbJpb6"."NilaiDbkbJpb6",NOW() FROM "DbkbJpb6" WHERE "DbkbJpb6"."TahunDbkbJpb6" = tempTahun2;

  INSERT INTO "DbkbJpb7" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb7","JenisDbkbJpb7","BintangDbkbJpb7","LantaiMinJpb7","LantaiMaxJpb7","NilaiDbkbJp75","CreatedAt") 
  SELECT "DbkbJpb7"."Provinsi_Kode","DbkbJpb7"."Daerah_Kode",TahunBaru,"DbkbJpb7"."JenisDbkbJpb7","DbkbJpb7"."BintangDbkbJpb7","DbkbJpb7"."LantaiMinJpb7","DbkbJpb7"."LantaiMaxJpb7","DbkbJpb7"."NilaiDbkbJp75",NOW() FROM "DbkbJpb7" WHERE "DbkbJpb7"."TahunDbkbJpb7" = tempTahun2;
  
  INSERT INTO "DbkbJpb8" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb8","LebarBentukMinDbkbJpb8","LebarBentukMaxDbkbJpb8","TinggiKolomMinDbkbJpb8","TinggiKolomMaxDbkbJpb8","NilaiDbkbJpb8","CreatedAt") 
  SELECT "DbkbJpb8"."Provinsi_Kode","DbkbJpb8"."Daerah_Kode",TahunBaru,"DbkbJpb8"."LebarBentukMinDbkbJpb8","DbkbJpb8"."LebarBentukMaxDbkbJpb8","DbkbJpb8"."TinggiKolomMinDbkbJpb8","DbkbJpb8"."TinggiKolomMaxDbkbJpb8","DbkbJpb8"."NilaiDbkbJpb8",NOW() FROM "DbkbJpb8" WHERE "DbkbJpb8"."TahunDbkbJpb8" = tempTahun2;
  
  INSERT INTO "DbkbJpb9" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb9","KelasDbkbJpb9","LantaiMinJpb9","LantaiMaxJpb9","NilaiDbkbJpb9","CreatedAt") 
  SELECT "DbkbJpb9"."Provinsi_Kode","DbkbJpb9"."Daerah_Kode",TahunBaru,"DbkbJpb9"."KelasDbkbJpb9","DbkbJpb9"."LantaiMinJpb9","DbkbJpb9"."LantaiMaxJpb9","DbkbJpb9"."NilaiDbkbJpb9",NOW() FROM "DbkbJpb9" WHERE "DbkbJpb9"."TahunDbkbJpb9" = tempTahun2;

  INSERT INTO "DbkbJpb13" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb13","KelasDbkbJpb13","LantaiMinJpb13","LantaiMaxJpb13","NilaiDbkbJpb13","CreatedAt") 
  SELECT "DbkbJpb13"."Provinsi_Kode","DbkbJpb13"."Daerah_Kode",TahunBaru,"DbkbJpb13"."KelasDbkbJpb13","DbkbJpb13"."LantaiMinJpb13","DbkbJpb13"."LantaiMaxJpb13","DbkbJpb13"."NilaiDbkbJpb13",NOW() FROM "DbkbJpb13" WHERE "DbkbJpb13"."TahunDbkbJpb13" = tempTahun2;
  
  INSERT INTO "DbkbJpb14" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb14","NilaiDbkbJpb14","CreatedAt") 
  SELECT "DbkbJpb14"."Provinsi_Kode","DbkbJpb14"."Daerah_Kode",TahunBaru,"DbkbJpb14"."NilaiDbkbJpb14",NOW() FROM "DbkbJpb14" WHERE "DbkbJpb14"."TahunDbkbJpb14" = tempTahun2;
  
  INSERT INTO "DbkbJpb15" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb15","JenisTangkiDbkbJpb15","KapasitasMinDbkbJpb15","KapasitasMaxDbkbJpb15","NilaiDbkbJpb15","CreatedAt") 
  SELECT "DbkbJpb15"."Provinsi_Kode","DbkbJpb15"."Daerah_Kode",TahunBaru,"DbkbJpb15"."JenisTangkiDbkbJpb15","DbkbJpb15"."KapasitasMinDbkbJpb15","DbkbJpb15"."KapasitasMaxDbkbJpb15","DbkbJpb15"."NilaiDbkbJpb15",NOW() FROM "DbkbJpb15" WHERE "DbkbJpb15"."TahunDbkbJpb15" = tempTahun2;
  
  INSERT INTO "DbkbJpb16" ("Provinsi_Kode","Daerah_Kode","TahunDbkbJpb16","KelasDbkbJpb16","LantaiMinJpb16","LantaiMaxJpb16","NilaiDbkbJpb16","CreatedAt") 
  SELECT "DbkbJpb16"."Provinsi_Kode","DbkbJpb16"."Daerah_Kode",TahunBaru,"DbkbJpb16"."KelasDbkbJpb16","DbkbJpb16"."LantaiMinJpb16","DbkbJpb16"."LantaiMaxJpb16","DbkbJpb16"."NilaiDbkbJpb16",NOW() FROM "DbkbJpb16" WHERE "DbkbJpb16"."TahunDbkbJpb16" = tempTahun2;
  
  INSERT INTO "DbkbDayaDukung" ("Provinsi_Kode","Daerah_Kode","ThnDbkbDayaDukung","TipeKonstruksi","NilaiDbkbDayaDukung","CreatedAt") 
  SELECT "DbkbDayaDukung"."Provinsi_Kode","DbkbDayaDukung"."Daerah_Kode",TahunBaru,"DbkbDayaDukung"."TipeKonstruksi","DbkbDayaDukung"."NilaiDbkbDayaDukung",NOW() FROM "DbkbDayaDukung" WHERE "DbkbDayaDukung"."ThnDbkbDayaDukung" = tempTahun2;
  
  INSERT INTO "DbkbMaterial" ("Provinsi_Kode","Daerah_Kode","Thn_Dbkb_Material","Pekerjaan_Kode","Kegiatan_Kode","NilaiDbkbMaterial","CreatedAt") 
  SELECT "DbkbMaterial"."Provinsi_Kode","DbkbMaterial"."Daerah_Kode",TahunBaru,"DbkbMaterial"."Pekerjaan_Kode","DbkbMaterial"."Kegiatan_Kode","DbkbMaterial"."NilaiDbkbMaterial",NOW() FROM "DbkbMaterial" WHERE "DbkbMaterial"."Thn_Dbkb_Material" = tempTahun2;
  
  INSERT INTO "DbkbMaterial" ("Provinsi_Kode","Daerah_Kode","Thn_Dbkb_Material","Pekerjaan_Kode","Kegiatan_Kode","NilaiDbkbMaterial","CreatedAt") 
  SELECT "DbkbMaterial"."Provinsi_Kode","DbkbMaterial"."Daerah_Kode",TahunBaru,"DbkbMaterial"."Pekerjaan_Kode","DbkbMaterial"."Kegiatan_Kode","DbkbMaterial"."NilaiDbkbMaterial",NOW() FROM "DbkbMaterial" WHERE "DbkbMaterial"."Thn_Dbkb_Material" = tempTahun2;
  
  INSERT INTO "DbkbMezanin" ("Provinsi_Kode","Daerah_Kode","TahunDbkbMezanin","NilaiDbkbMezanin","CreatedAt") 
  SELECT "DbkbMezanin"."Provinsi_Kode","DbkbMezanin"."Daerah_Kode",TahunBaru,"DbkbMezanin"."NilaiDbkbMezanin",NOW() FROM "DbkbMezanin" WHERE "DbkbMezanin"."TahunDbkbMezanin" = tempTahun2;
  
  INSERT INTO "DataNir" ("Provinsi_Kode","Daerah_Kode","Kecamatan_Kode","Kelurahan_Kode","Tahun","NomerDokumen","Kpbb_Id","Kanwil_Id","JenisDokumen","Znt_Kode","Nir","CreatedAt") 
  SELECT "DataNir"."Provinsi_Kode","DataNir"."Daerah_Kode","DataNir"."Kecamatan_Kode","DataNir"."Kelurahan_Kode",TahunBaru,"DataNir"."NomerDokumen","DataNir"."Kpbb_Id","DataNir"."Kanwil_Id","DataNir"."JenisDokumen","DataNir"."Znt_Kode","DataNir"."Nir",NOW() FROM "DataNir" WHERE "DataNir"."Tahun" = tempTahun2;
  
  INSERT INTO "FasDepJpbKlsBintang" ("Provinsi_Kode","Daerah_Kode","TahunDep","Fasilitas_Kode","Jpb_Kode","KlsBintang","NilaiFasilitasKlsBintang","CreatedAt") 
  SELECT "FasDepJpbKlsBintang"."Provinsi_Kode","FasDepJpbKlsBintang"."Daerah_Kode",TahunBaru,"FasDepJpbKlsBintang"."Fasilitas_Kode","FasDepJpbKlsBintang"."Jpb_Kode","FasDepJpbKlsBintang"."KlsBintang","FasDepJpbKlsBintang"."NilaiFasilitasKlsBintang",NOW() FROM "FasDepJpbKlsBintang" WHERE "FasDepJpbKlsBintang"."TahunDep" = tempTahun2;
  
  INSERT INTO "FasDepMinMax" ("Provinsi_Kode","Daerah_Kode","TahunDepMinMax","Fasilitas_Kode","KlsDepMin","KlsDepMax","NilaiDepMinMax","CreatedAt") 
  SELECT "FasDepMinMax"."Provinsi_Kode","FasDepMinMax"."Daerah_Kode",TahunBaru,"FasDepMinMax"."Fasilitas_Kode","FasDepMinMax"."KlsDepMin","FasDepMinMax"."KlsDepMax","FasDepMinMax"."NilaiDepMinMax",NOW() FROM "FasDepMinMax" WHERE "FasDepMinMax"."TahunDepMinMax" = tempTahun2;
  
  INSERT INTO "FasNonDep" ("Provinsi_Kode","Daerah_Kode","Thn_Non_Dep","Fasilitas_Kode","CreatedAt") 
  SELECT "FasNonDep"."Provinsi_Kode","FasNonDep"."Daerah_Kode",TahunBaru,"FasNonDep"."Fasilitas_Kode",NOW() FROM "FasNonDep" WHERE "FasNonDep"."Thn_Non_Dep" = tempTahun2;
  
  INSERT INTO "KayuUlin" ("Provinsi_Kode","Daerah_Kode","ThnStatusKayuUlin","StatusKayuUlin","CreatedAt") 
  SELECT "KayuUlin"."Provinsi_Kode","KayuUlin"."Daerah_Kode",TahunBaru,"KayuUlin"."StatusKayuUlin",NOW() FROM "KayuUlin" WHERE "KayuUlin"."ThnStatusKayuUlin" = tempTahun2;
  
  INSERT INTO "TempatPembayaranSPPTMasal" ("Provinsi_Kode","Dati2_Kode","Tahun","Kecamatan_Kode","Kelurahan_Kode","Kanwil_Kode","KPPBB_Kode","BankTunggal_Kode","BankPersepsi_Kode","TP_Kode","CreatedAt") 
  SELECT "TempatPembayaranSPPTMasal"."Provinsi_Kode","TempatPembayaranSPPTMasal"."Dati2_Kode",TahunBaru,"TempatPembayaranSPPTMasal"."Kecamatan_Kode","TempatPembayaranSPPTMasal"."Kelurahan_Kode","TempatPembayaranSPPTMasal"."Kanwil_Kode","TempatPembayaranSPPTMasal"."KPPBB_Kode","TempatPembayaranSPPTMasal"."BankTunggal_Kode","TempatPembayaranSPPTMasal"."BankPersepsi_Kode","TempatPembayaranSPPTMasal"."TP_Kode",NOW() FROM "TempatPembayaranSPPTMasal" WHERE "TempatPembayaranSPPTMasal"."Tahun" = tempTahun2;

  return TahunBaru;

end;$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
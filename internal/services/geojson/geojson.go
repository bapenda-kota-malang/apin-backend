package geojson

import (
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"
	gj "github.com/paulmach/go.geojson"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/geojson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "geojson"

func Create(input m.CreateDto) (any, error) {
	var data m.GeoJson

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.GeoJson
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.GeoJson{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetListGeoJson(input m.FilterDto) (any, error) {
	var datas []m.GeoJson
	var count int64
	var lists [][]byte

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.GeoJson{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&datas)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", datas)
	}

	var list []byte
	for _, data := range datas {
		list = setGeoJson(data)
		lists = append(lists, list)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: lists,
	}, nil
}

func GetDetail(id int) (any, error) {
	var data *m.GeoJson

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func setGeoJson(data m.GeoJson) []byte {

	featureCollection := gj.NewFeatureCollection()
	feature := gj.NewMultiPolygonFeature(data.GeoJsonPosition())

	feature.SetProperty("FID_", data.FID_)
	feature.SetProperty("Entity", data.Entity)
	feature.SetProperty("Layer", data.Layer)
	feature.SetProperty("Color", data.Color)
	feature.SetProperty("Linetype", data.Linetype)
	feature.SetProperty("Elevation", data.Elevation)
	feature.SetProperty("LineWt", data.LineWt)
	feature.SetProperty("RefName", data.RefName)
	feature.SetProperty("d_nop", data.D_Nop)
	feature.SetProperty("d_luas", data.D_Luas)
	feature.SetProperty("NAMA", data.Nama)
	feature.SetProperty("NO", data.No)
	feature.SetProperty("ALAMAT", data.Alamat)
	feature.SetProperty("RT", data.Rt)
	feature.SetProperty("RW", data.Rw)
	feature.SetProperty("KELURAHAN", data.Kelurahan)
	feature.SetProperty("KECAMATAN", data.Kecamatan)
	feature.SetProperty("KOTA", data.Kota)
	feature.SetProperty("LUAS_TANAH", data.Luas_tanah)
	feature.SetProperty("KODE_ZNT", data.Kode_ZNT)
	feature.SetProperty("JNIS_TANAH", data.Jnis_Tanah)
	feature.SetProperty("J_BANGUNAN", data.J_Bangunan)
	feature.SetProperty("L_BANGUN", data.L_Bangun)
	feature.SetProperty("NJOP_BUMI", data.Njop_Bumi)
	feature.SetProperty("NJOP_BGN", data.Njop_Bgn)
	feature.SetProperty("NJOP_TOTAL", data.Njop_Total)
	feature.SetProperty("PENDATAAN", data.Pendataan)
	feature.SetProperty("KETERANGAN", data.Keterangan)
	feature.SetProperty("EDIT_KAV", data.Edit_Kav)
	feature.SetProperty("FUNGSI", data.Fungsi)
	feature.SetProperty("JNIS_PAJAK", data.Jnis_Pajak)

	featureCollection.AddFeature(feature)
	result, err := featureCollection.MarshalJSON()
	if err != nil {
		return nil
	}

	return result
}

func GetDetailGeoJson(id int) (any, error) {
	var data *m.GeoJson

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	fresult := setGeoJson(*data)

	return rp.OKSimple{
		Data: fresult,
	}, nil
}

func Update(id int, input m.UpdateDto) (any, error) {
	var data *m.GeoJson
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {
	var data *m.GeoJson
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}

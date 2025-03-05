package row

import (
	"block1_http/handler"
	"block1_http/signals"
	"encoding/json"
	"fmt"
)

func Get_last_row(adress string) []byte {
	db := handler.Open_db(adress) // открываем базу данных
	defer db.Close()              // после всех манипуляций базу данных закрываем

	signals := signals.Signals{}
	row_last := db.QueryRow("SELECT * FROM block1 ORDER BY id DESC LIMIT 1")
	err := row_last.Scan(
		&signals.Id,
		&signals.Date,
		&signals.O2_slev,
		&signals.O2_sprav,
		&signals.Q_gaz,
		&signals.T_para_nitk_a,
		&signals.T_para_nitk_b,
		&signals.P_para,
		&signals.Q_per_par,
		&signals.Q_pit_voda,
		&signals.T_vozdh_posl_rvpa,
		&signals.T_vozdh_posl_rvpb,
		&signals.T_vozdh_posl_tvp_slev,
		&signals.T_vozdh_posl_tvp_sprav,
		&signals.T_posle_ekonomiz_slev,
		&signals.T_posle_ekonomiz_sprav,
		&signals.T_dym_zarvp_slev,
		&signals.T_dym_zarvp_sprav,
		&signals.T_dym_zatvp_slev,
		&signals.T_dym_zatvp_sprav,
		&signals.T_pitvod_kotl,
		&signals.T_uhodgaz_ventur1,
		&signals.T_uhodgaz_ventur2,
		&signals.T_mazut,
		&signals.T_uhodgaz_ventur3,
		&signals.T_uhodgaz_ventur4,
		&signals.T_kontensat_sbor,
		&signals.T_kontensat_vihkollect,
		&signals.T_vod_kollect_slev,
		&signals.T_vod_kollect_sprav,
		&signals.T_gaz_dsa,
		&signals.T_gaz_dsb,
		&signals.T_vozdh_pered_rvpa,
		&signals.T_vozdh_pered_rvpb,
		&signals.T_vozdh_pered_tvpslev,
		&signals.T_vozdh_pered_tvpsprav,
		&signals.H_baraban_slev,
		&signals.H_baraban_sprav,
		&signals.P_par_do_sk,
		&signals.Q_par,
		&signals.T_par_posle_sk,
		&signals.Vakum,
		&signals.T_par_psg1_prov1,
		&signals.T_par_psg1_prov2,
		&signals.T_par_psg2,
		&signals.T_vod_do_psg1,
		&signals.T_vod_posle_psg1,
		&signals.T_vod_vihod_psg12,
		&signals.T_par_posle_ou1,
		&signals.T_par_posle_ou2,
		&signals.T_kondesat_do_kn,
		&signals.T_cirk_vod_posle_kondensat,
		&signals.T_cirk_vod_do_kondensat,
		&signals.T_par_cnd_sprav,
		&signals.T_par_cnd_slev,
		&signals.T_par_uplotn_kollect,
		&signals.T_kondest_posle_pnd4,
		&signals.T_vod_obvod_pvd7,
		&signals.T_par_pered_pn130,
		&signals.T_vod_posle_pvd7,
		&signals.P_par_k_psg1,
		&signals.P_par_k_psg2,
		&signals.Akt_stal_vozb_paz_7,
		&signals.Akt_stal_vozb_paz_21,
		&signals.Akt_stal_vozb_paz_36,
		&signals.Akt_stal_vozb_paz_50,
		&signals.Akt_stal_vozb_paz_64,
		&signals.Akt_stal_vozb_paz_77,
		&signals.Akt_stal_vozb_paz_3g7,
		&signals.Akt_stal_vozb_paz_25g10,
		&signals.Akt_stal_vozb_paz_44g12,
		&signals.Akt_stal_vozb_paz_3g13,
		&signals.Akt_stal_vozb_paz_25g16,
		&signals.Akt_stal_vozb_paz_44g18,
		&signals.T_vod_sobstv_nuzht,
		&signals.T_vod_glav_korp_truba1,
		&signals.T_vod_glav_korp_truba2,
		&signals.T_vod_za_psn1,
		&signals.T_vod_za_psn2)
	if err != nil {
		fmt.Println(err)
	}

	json_byte, err := json.Marshal(signals)
	if err != nil {
		fmt.Println(err)
	}
	return json_byte

}

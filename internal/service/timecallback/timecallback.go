package timecallback

import (
	"oauth2/internal/core"
	"oauth2/internal/model"
	"time"

	"github.com/xxjwxc/public/mylog"
)

// TimeCallBackToken 时间回调
func TimeCallBackToken() {
	orm := core.Dao.GetDBw()

	//删除access_token
	err := model.AccessTokenTblMgr(orm.DB).Where("expires <= ?", time.Now()).Delete(&model.AccessTokenTbl{}).Error
	if err != nil {
		mylog.Error(err)
	}

	//删除refresh_token
	err = model.RefreshTokenTblMgr(orm.DB).Where("expires <= ?", time.Now()).Delete(&model.RefreshTokenTbl{}).Error
	if err != nil {
		mylog.Error(err)
	}
}

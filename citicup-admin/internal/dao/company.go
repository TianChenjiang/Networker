package dao

import (
	"citicup-admin/internal/model"
)

var company_e interface{} = &model.Company{}

//查询所有的公司
func (d *Dao) GetAllCompanies() (companies []*model.Company, err error) {
	d.db.Model(company_e).Find(&companies)
	return
}

//根据公司Id查询
func (d *Dao) GetCompanyById(id uint) (company model.Company, err error) {
	d.db.Model(company_e).Where(&model.Company{
		ID: id,
	}).First(&company)
	return
}

//更新公司信息
func (d *Dao) UpdateCompany(entity *model.Company) (company model.Company, err error) {
	d.db.Model(company_e).Where(&model.Company{
		ID: entity.ID,
	}).Update(entity).Find(&company)
	return
}

//删除公司
func (d *Dao) DeleteCompany(id uint) (company model.Company, err error) {
	d.db.Model(company_e).Where(&model.Company{
		ID: id,
	}).Delete(&company)
	return
}

func (d *Dao) InsertCompany(entity *model.Company) (newCompany model.Company, err error) {
	d.db.Model(company_e).Create(entity).Find(&newCompany)
	return
}

func (d *Dao) GetMarket(companyID uint) (market model.Market, err error) {
	company, _ := d.GetCompanyById(companyID)
	d.db.Preload("Market")
	d.db.Model(market_e).Where("id = ?", company.MarketID).Find(&market)
	return
}

func (d *Dao) GetCompanyBySymbol(symbol string) (company model.Company, err error) {
	d.db.Model(company_e).Where(&model.Company{
		Symbol: symbol,
	}).Find(&company)
	return
}
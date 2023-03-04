package repositories

import (
	"database/sql"
	"github.com/A-walker-ninght/mini-seckill/common"
	_interface "github.com/A-walker-ninght/mini-seckill/interface"
	"github.com/A-walker-ninght/mini-seckill/models"
	"strconv"
)

type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) _interface.Repository {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
	}
}

func (p *ProductManager) Conn() error {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return nil
}

func (p *ProductManager) Insert(m models.Model) (int64, error) {
	if err := p.Conn(); err != nil {
		return 0, err
	}
	data := m.(*models.Product)

	sql := "INSERT " + p.table + " SET productName = ?, productNum = ?, productImage = ?, productUrl = ?"
	stmt, errSql := p.mysqlConn.Prepare(sql)
	if errSql != nil {
		return 0, errSql
	}
	result, errStmt := stmt.Exec(data.ProductName, data.ProductNum, data.ProductImage, data.ProductUrl)
	if errStmt != nil {
		return 0, errStmt
	}
	return result.LastInsertId()
}

func (p *ProductManager) Delete(i int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}
	sql := "delete from " + p.table + " where ID = ?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	_, err = stmt.Exec(strconv.FormatInt(i, 10))
	if err != nil {
		return false
	}
	return true
}

func (p *ProductManager) Update(m models.Model) error {
	if err := p.Conn(); err != nil {
		return err
	}
	sql := "update " + p.table + " set productName=?,productNum=?,productImage=?,productUrl=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	data := m.(*models.Product)
	_, err = stmt.Exec(data.ProductName, data.ProductNum, data.ProductImage, data.ProductUrl)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductManager) SelectByKey(i int64) (models.Model, error) {
	if err := p.Conn(); err != nil {
		return &models.Product{}, err
	}
	sql := "SELECT * FROM " + p.table + " where ID = " + strconv.FormatInt(i, 10)
	rows, err := p.mysqlConn.Query(sql)
	if err != nil {
		return &models.Product{}, err
	}

	result := common.GetResultRow(rows) // result为空？？？？
	if len(result) == 0 {
		return &models.Product{}, nil
	}
	product := &models.Product{} // 注意传入指针
	common.DataToStructByTagSql(result, product)
	return product, nil
}

func (p *ProductManager) SelectAll() ([]models.Model, error) {
	if err := p.Conn(); err != nil {
		return nil, err
	}
	sql := "select * from " + p.table

	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}

	products := make([]models.Model, 0)
	for _, v := range result {
		product := &models.Product{}
		common.DataToStructByTagSql(v, product)
		products = append(products, product)
	}

	return products, nil
}

func (p *ProductManager) SelectAllWithInfo() (map[int]map[string]string, error) {
	return nil, nil
}

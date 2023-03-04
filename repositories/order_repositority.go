package repositories

import (
	"database/sql"
	"github.com/A-walker-ninght/mini-seckill/common"
	_interface "github.com/A-walker-ninght/mini-seckill/interface"
	"github.com/A-walker-ninght/mini-seckill/models"
	"strconv"
)

type OrderManagerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func NewOrderManagerRepository(table string, db *sql.DB) _interface.Repository {
	return &OrderManagerRepository{
		table:     table,
		mysqlConn: db,
	}
}

func (o *OrderManagerRepository) Conn() error {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

func (o *OrderManagerRepository) Insert(m models.Model) (int64, error) {
	if err := o.Conn(); err != nil {
		return 0, err
	}
	sql := "INSERT " + o.table + " SET userID = ?, productID = ?, orderStatus = ?"
	stmt, errSql := o.mysqlConn.Prepare(sql)
	if errSql != nil {
		return 0, errSql
	}
	order := m.(*models.Order)
	result, errStmt := stmt.Exec(order.UserID, order.ProductId, order.OrderStatus)
	if errStmt != nil {
		return 0, errStmt
	}
	return result.LastInsertId()
}

func (o *OrderManagerRepository) Delete(i int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}
	sql := "delete from " + o.table + " where ID = ?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	_, err = stmt.Exec(strconv.FormatInt(i, 10))
	if err != nil {
		return false
	}
	return true
}

func (o *OrderManagerRepository) Update(m models.Model) error {
	if err := o.Conn(); err != nil {
		return err
	}
	sql := "update " + o.table + " set userID=?,productID=?,orderStatus=?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	order := m.(*models.Order)
	_, err = stmt.Exec(order.UserID, order.ProductId, order.OrderStatus)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderManagerRepository) SelectByKey(i int64) (models.Model, error) {
	if err := o.Conn(); err != nil {
		return &models.Order{}, err
	}
	sql := "SELECT * FROM " + o.table + " where ID = " + strconv.FormatInt(i, 10)
	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		return &models.Order{}, err
	}

	result := common.GetResultRow(rows) // result为空？？？？
	if len(result) == 0 {
		return &models.Order{}, nil
	}
	ord := &models.Order{} // 注意传入指针
	common.DataToStructByTagSql(result, ord)
	return ord, nil
}

func (o *OrderManagerRepository) SelectAll() ([]models.Model, error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}
	sql := "select * from " + o.table

	rows, err := o.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}

	orders := make([]models.Model, 0)
	for _, v := range result {
		ord := &models.Order{}
		common.DataToStructByTagSql(v, ord)
		orders = append(orders, ord)
	}

	return orders, nil
}

func (o *OrderManagerRepository) SelectAllWithInfo() (map[int]map[string]string, error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}
	sql := "select o.ID, p.productName, o.orderStatus from seckill.order as o, seckill.product as p where o.productID = p.ID"
	rows, err := o.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

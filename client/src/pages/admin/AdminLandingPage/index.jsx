import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import { adminGetUsers, adminDeleteUser } from "../../../store/Actions/adminActions";
import HeaderAdmin from "../../../pages/menu/HeaderAdmin";
import Loading from "../../../components/Loading";
import Footer from "../../../components/Footer";
// import Footer from "../../../components/Footer";
import "./style.css";
import { Layout, Table, Button, Divider, Popconfirm, message } from "antd";
const { Content } = Layout;

class AdminLandingPage extends Component {
  constructor(props) {
    super(props);
    this.columns = [
      {
        title: "Employee ID",
        dataIndex: "employee_number",
        key: "employee_number",
        width: "10%"
      },
      {
        title: "Name",
        dataIndex: "name",
        key: "name",
        width: "20%"
      },
      {
        title: "Email",
        dataIndex: "email",
        key: "email",
        width: "15%"
      },
      {
        title: "Position",
        dataIndex: "position",
        key: "position",
        width: "15%"
      },
      {
        title: "Role",
        dataIndex: "role",
        key: "role",
        width: "10%"
      },
      {
        title: "Action",
        key: "action",
        width: "25%",
        render: (text, record) => (
          //  (record.employee_number))
          <span>
            <Button
              onClick={() => {
                this.editUser(this.props.users, record.employee_number);
              }}
              type="primary"
            >
              Edit
            </Button>

            <Divider type="vertical" />
            <Popconfirm
              placement="top"
              title={"Are you sure delete this employee?"}
              onConfirm={() => {
                this.adminDeleteUser(this.props.users, record.employee_number);
                message.success("Employee has been delete!");
              }}
              okText="Yes"
              cancelText="No"
            >
              <Button type="danger">Delete</Button>
            </Popconfirm>
            <Divider type="vertical" />
            {record.role === "supervisor" || record.role === "employee" ? (
              <Button
                onClick={() => {
                  this.editBalance(record.employee_number);
                }}
                type="primary"
              >
                Edit Leave Balance
              </Button>
            ) : (
                ""
              )}
          </span>
        )
      }
    ];
  }

  componentWillMount() {
    console.log(" ----------------- Admin-Page ----------------- ");
  }

  componentDidMount() {
    if (localStorage.getItem("role") !== "admin") {
      this.props.history.push("/");
    }
    this.props.adminGetUsers();
  }

  editUser = (users, id) => {
    this.props.history.push({
      pathname: "/admin/edit-user/" + id,
      state: { users: users }
    });
  };

  editBalance = id => {
    this.props.history.push({
      pathname: "/admin/edit-balance/" + id
    });
  };

  adminDeleteUser = (users, id) => {
    this.props.adminDeleteUser(users, id);
  };

  onShowSizeChange(current, pageSize) {
    console.log(current, pageSize);
  }

  render() {
    if (this.props.loading) {
      return <Loading />;
    } else {
      return (
        <div>
          <Layout className="layout" >
            <HeaderAdmin />
            <Content
              className="container"
              style={{
                display: "flex",
                margin: "20px 15px 0",
                justifyContent: "center",
                paddingBottom: 70,
              }}
            >
              <div style={{ padding: "20px", backgroundColor: "#fff", width: "1700px" }}>
                <Table
                  columns={this.columns}
                  dataSource={this.props.users}
                  rowKey={record => record.employee_number}
                  pagination={{
                    className: "my-pagination",
                    defaultCurrent: 1,
                    defaultPageSize: 10,
                    total: `${this.props.users && this.props.users.length}`,
                    // showSizeChanger: this.onShowSizeChange
                  }}
                />
              </div>
            </Content>
            <Footer />
          </Layout>
        </div>
      );
    }
  }
}

const mapStateToProps = state => ({
  loading: state.adminReducer.loading,
  users: state.adminReducer.users
});

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      adminGetUsers,
      adminDeleteUser
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AdminLandingPage);

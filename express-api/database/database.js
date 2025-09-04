import mysql from 'mysql';
const con = mysql.createConnection({
  host: 'localhost',
  user: 'appuser',
  password: 'apppassword',
  database: 'myapp'
});
con.connect();
export default con;
import express from 'express';
import bodyParser from 'body-parser';
import userRouter from './routes/userRoute.js';
import dotenv from "dotenv";

dotenv.config({path: '../.env'});

const app = express();
const port = 3000;
app.use(bodyParser.json());
app.use('/api', userRouter);

app.use((req, res) => {
  res.status(404).send('404 - Page Not Found');
});

app.listen(port, () => {
    console.log(`app running in port ${port}`)
})
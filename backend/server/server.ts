import express = require("express");
import cors from "cors";
import * as dotenv from "dotenv";
import { errorHandler, notFoundHandler } from "../app/handlers/Exceptions";
import routes from "../app/routes";
// import 'module-alias/register';

dotenv.config();

const app: express.Application = express();
const port: number = parseInt(process.env.PORT as string, 10) || 3000;

app.use(cors());
app.use(express.json());

app.use(routes);

app.use(errorHandler);
app.use(notFoundHandler);

app.listen(port, function () {
    console.log(`Listening on port ${port}`);
});
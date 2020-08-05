import express from 'express';
import * as controllers from './controllers';
import * as validators from './validators'
import * as middlewares from './middlewares'

const routes = express.Router(); 

// const examples = "/examples";
routes.get('/api', controllers.SvgController.get);
// routes.get(`${examples}`, [middlewares.ExampleMiddleware, validators.ExamplesValidator], controllers.ExamplesController.index);
// routes.get(`${examples}/:id`, [middlewares.ExampleMiddleware, middlewares.ExampleMiddleware], ExampleController.show);
// routes.post(`${examples}`, [middlewares.ExampleMiddleware, middlewares.ExampleMiddleware], ExampleController.store);
// routes.put(`${examples}/:id`, [middlewares.ExampleMiddleware, middlewares.ExampleMiddleware], ExampleController.update);
// routes.delete(`${examples}/:id`, [middlewares.ExampleMiddleware, middlewares.ExampleMiddleware], ExampleController.destroy);

export default routes;

// HACKED!

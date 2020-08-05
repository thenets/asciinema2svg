import { Request, Response, NextFunction } from 'express';

export default async function(request: Request, response: Response, next: NextFunction) {
  console.log('ExampleMiddleware');

  return next();
}
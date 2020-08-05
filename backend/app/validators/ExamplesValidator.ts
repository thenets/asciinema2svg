import { celebrate, Joi } from 'celebrate';
import { Request, Response, NextFunction } from 'express';

export default async function(request: Request, response: Response, next: NextFunction){
  console.log('ExamplesValidator');
  
  // celebrate({
  //     body: Joi.object().keys({
  //         title: Joi.string().required(),
  //         link: Joi.string().required(),
  //         description: Joi.string().required(),
  //         tags: Joi.string().required(),
  //     })
  // },
  // {
  //     abortEarly: false
  // })

  return next();
}
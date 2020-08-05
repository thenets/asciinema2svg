import { Request, Response } from 'express';
// import Example from '../models/Example';
import fs from 'fs';
import {promisify} from 'util'
import {render, SvgTerm} from 'svg-term'
import path from 'path'

const readFile = promisify(fs.readFile);

export async function get(request: Request, response: Response){
  const basePath = path.resolve('public', 'ascii.cast');
  const data = String(await readFile(basePath));

  const svg = render(data, render);
  return response.send(svg);
  // return response.json(data);
}
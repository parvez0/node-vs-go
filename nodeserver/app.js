const express = require('express');
const logger = require('morgan');

const app = express();

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

const sleep = (time) => {
    return new Promise(resolve => {
        setTimeout( () => { resolve(); }, time);
    });
}

app.get('/', (req, res) => {
   res.status(200).json({ 'success': true, data: { message: 'server working !!!' } });
});

app.post('/', async (req, res) => {
    const time = req.body.sleepTime || 10000;
    await sleep(time);
    res.status(200).json({ 'success': true, data: { message: `responding after ${time} milliseconds.` } });
});

const port = process.env.PORT || '3000';
app.listen(port, () => { console.log('server is listening on port: ', port) });
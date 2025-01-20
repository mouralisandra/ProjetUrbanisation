const express = require('express');
const app = express();
const PORT = 3000;

// Dummy endpoint to fetch a sales report
app.get('/report/sales', (req, res) => {
    res.json({
        report: 'Monthly Sales Report',
        totalSales: 10000,
        generatedAt: new Date().toISOString()
    });
});

// Dummy endpoint to fetch summary
app.get('/report/summary', (req, res) => {
    res.json({
        summary: 'Quarterly Sales Summary',
        totalProfit: 5000,
        generatedAt: new Date().toISOString()
    });
});

// Dummy endpoint to fetch summary
app.post('/report/missing', (req, res) => {
    res.json({
        summary: 'Reported luggage missing',
        generatedAt: new Date().toISOString()
    });
});

app.listen(PORT, () => {
    console.log(`Sales Reporting Service is running on port ${PORT}`);
});

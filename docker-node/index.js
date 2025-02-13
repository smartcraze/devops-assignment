const express = require("express");

express.get("/", (req, res) => {
    res.json({
        message: "Hello world"
    })
})

express.listen(3000);
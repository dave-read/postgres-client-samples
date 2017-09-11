
var urlencode = require('urlencode');

const delimeters = "!#$&()*+,/:;=?@[]"
const minimal = "Aaaaaaaa"

const passwords = [
    "Passw!rd",
    "Passw#rd",
    "Passw$rd",
    "Passw&rd",
    "Passw(rd",
    "Passw)rd",
    "Passw*rd",
    "Passw+rd",
    "Passw,rd",
    "Passw/rd",
    "Passw:rd",
    "Passw;rd",
    "Passw=rd",
    "Passw?rd",
    "Passw@rd",
    "Passw[rd",
    "Passw]rd"
]
passwords.forEach(function(password){
    console.log("password           ->"+password)
    console.log("encodeURI          ->"+encodeURI(password))
    console.log("urlencode          ->"+urlencode(password))
    console.log("encodeURIComponent ->"+encodeURIComponent(password))
    console.log("==============================================================")
    
});
process.exit(1)
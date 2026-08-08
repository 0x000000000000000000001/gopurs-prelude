const fs = require('fs');
const fn = JSON.parse(fs.readFileSync('output/Test.Main/corefn.json', 'utf8'));
const decl = fn.decls.find(d => d.identifier === 'testSignum');
console.log(JSON.stringify(decl, null, 2));

import { Command } from 'commander';


const program = new Command();

program
    .option('-q, --query <string>', 'The query')
    .option('-d, --data-source <string>', 'The data source')
    .parse(process.argv);

const options = program.opts();
if (typeof options.query !== "string") {
    throw Error("The query is not defined");
}

if (typeof options.dataSource !== "string") {
    throw Error("The query is not defined");
}

const result = [
    { "value": "в", "nextNode": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/в", "operator": "https://w3id.org/tree#SubstringRelation", "node": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/" },
    { "value": "3", "nextNode": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/3", "operator": "https://w3id.org/tree#SubstringRelation", "node": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/" },
    { "value": "4", "nextNode": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/4", "operator": "https://w3id.org/tree#SubstringRelation", "node": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/" },
    { "value": "a", "nextNode": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/д", "operator": "https://w3id.org/tree#SubstringRelation", "node": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/" },
    { "value": "5", "nextNode": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/5", "operator": "https://w3id.org/tree#SubstringRelation", "node": "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/" }
];

console.log(JSON.stringify(result));

import { Command } from 'commander';


const program = new Command();
/**
 * Hacky way to avoid the error from golang 
 */
console.warn = () => {};
console.error = () => {};

program
    .option('-l, --query <string>', 'The query')
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

const compenentJsError = `2023-06-23T07:09:43.951Z [Components.js] warn: Detected deprecated context URL 'https://linkedsoftwaredependencies.org/bundles/npm/componentsjs/^4.0.0/components/context.jsonld' in /home/id357/Documents/PhD/coding/comunica_filter_benchmark/node_modules/@treecg/connector-types/components/components.jsonld. Prefer using version '^5.0.0' instead.
2023-06-23T07:09:45.969Z [Components.js] warn: Detected deprecated context URL 'https://linkedsoftwaredependencies.org/bundles/npm/componentsjs/^4.0.0/components/context.jsonld' in /home/id357/Documents/PhD/coding/comunica_filter_benchmark/node_modules/@treecg/connector-types/components/types.jsonld. Prefer using version '^5.0.0' instead.`;
console.error(compenentJsError);
console.log(JSON.stringify(result));

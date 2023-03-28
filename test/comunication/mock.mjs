import { Command } from 'commander';

program
    .requiredOption('-q, --query <string>', 'The query')
    .requiredOption('-l, --lenient <number>', 'Lenient option')
    .parse(process.argv);

const options = program.opts();
if (options.length === 0) {
    process.exit(1)
}

if (options.lenient === false) {
    process.exit(1)
}

process.stdout.write(`
[
    {"value":"в","nextNode":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/в","operator":"https://w3id.org/tree#SubstringRelation","node":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/"},
    {"value":"3","nextNode":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/3","operator":"https://w3id.org/tree#SubstringRelation","node":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/"},
    {"value":"4","nextNode":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/4","operator":"https://w3id.org/tree#SubstringRelation","node":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/"},
    {"value":"a","nextNode":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/д","operator":"https://w3id.org/tree#SubstringRelation","node":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/"},
    {"value":"5","nextNode":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/5","operator":"https://w3id.org/tree#SubstringRelation","node":"https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/"}
]\n
        `
)
import { Command } from 'commander';
import { QueryEngineFactory } from "@comunica/query-sparql-link-traversal";
import { KeysExtractLinksTree } from '@comunica/context-entries-link-traversal';

const program = new Command();
const default_query = (limit) => `
PREFIX tree: <https://w3id.org/tree#>
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>

SELECT ?node ?nextNode ?operator ?value WHERE {
  ?node tree:relation ?relation .
  ?relation tree:node ?nextNode .
  
  ?relation rdf:type ?operator.
  ?relation tree:value ?value .
} LIMIT ${limit}`;
program
    .requiredOption('-l, --limit <number>', 'The limit of the query', Number.MAX_SAFE_INTEGER)
    .requiredOption('-d, --data-source <string>', 'The data source', 'http://localhost:3000/ldes/test')
    .parse(process.argv);

const options = program.opts();
const query = default_query(options.limit);

const datasource = options.dataSource;
const config = "./comunica-js/config.json";

console.log(`registering the query of ${datasource} data source with limit ${options.limit}`)

function getRelations() {

    return new Promise(async (resolve, _reject) => {
        const engine = await new QueryEngineFactory().create({ configPath: config });
        const bindingsStream = await engine.queryBindings(query,
            {
                sources: [datasource],
                lenient: true,
                [KeysExtractLinksTree.strictTraversal.name]: false,
            });
        bindingsStream.on('data', (binding) => {
            console.log(JSON.stringify(
                {
                    "operator": binding.get('operator').value,
                    "value": binding.get('value').value,
                    "nextNode": binding.get('nextNode').value,
                    "node": binding.get('node').value
                }
            )
            );
        });

        bindingsStream.on('error', (_error) => {
            resolve();
        });

        bindingsStream.on('end', () => {
            resolve()
        });
    });


}


await getRelations();
process.exit(0);
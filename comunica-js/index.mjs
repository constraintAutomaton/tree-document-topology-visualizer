import { Command } from 'commander';
import { QueryEngineFactory } from "@comunica/query-sparql-link-traversal";
import { KeysExtractLinksTree } from '@comunica/context-entries-link-traversal';

const program = new Command();

program
    .requiredOption('-q, --query [string...]', 'The query')
    .requiredOption('-d, --data-source <string>', 'The data source')
    .parse(process.argv);

const options = program.opts();

const query = options.query.join(' ');
const datasource = options.dataSource;
const config = "./comunica-js/config.json";

function getRelations() {
    return new Promise(async (resolve, _reject) => {
        const result = [];
        const engine = await new QueryEngineFactory().create({ configPath: config });
        const bindingsStream = await engine.queryBindings(query,
            {
                sources: [datasource],
                lenient: true,
                [KeysExtractLinksTree.strictTraversal.name]: false,
            });
        bindingsStream.on('data', (binding) => {
            result.push(
                {
                    "operator": binding.get('operator').value,
                    "value": binding.get('value').value,
                    "nextNode": binding.get('nextNode').value,
                    "node": binding.get('node').value
                }
            )
        });

        bindingsStream.on('error', (_error) => {
            resolve(result);
        });

        bindingsStream.on('end', () => {
            resolve(result)
        });
    });


}

const result = await getRelations();
console.log(JSON.stringify(result));
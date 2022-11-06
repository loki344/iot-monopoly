export default function (plop) {
	plop.setGenerator('component', {
		description: 'Create a new component',
		prompts: [
			{
				// Name your Component
				type: 'input',
				name: 'name',
				message: 'What is the name of your component?'
			},
			{
				// Atomic design category
				type: 'input',
				name: 'category',
				message: 'What is it? Atom, Molecule, Organism, Layout, Template, Page, Wrapper'
			}
		],
		actions: [
			{
				// Create the component files
				type: 'addMany',
				destination: './src/lib/components/{{lowerCase category}}',
				base: `.templates/component/definition`,
				templateFiles: `.templates/component/definition/*.hbs`
			},
			{
				// Create a test suite for the component
				type: 'add',
				path: './tests/{{properCase name}}.spec.ts',
				templateFile: '.templates/component/tests/test.spec.ts.hbs'
			},
			{
				// Create the story files
				type: 'addMany',
				destination: './src/lib/components/{{lowerCase category}}/story',
				base: `.templates/component/histoire`,
				templateFiles: `.templates/component/histoire/*.hbs`
			}
		]
	});
}

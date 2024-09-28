import { defineConfig } from 'vite';
import path from 'path';
import dts from 'vite-plugin-dts';

export default defineConfig({
	build: {
		lib: {
			entry: path.resolve(__dirname, 'src/index.ts'),
			name: 'Common',
			fileName: (format) => `index.${format}.js`,
			// formats: ['es', 'umd'],
		},
		rollupOptions: {
			external: [],
			output: {
				globals: {},
			},
		},
		outDir: 'dist',
	},
	plugins: [dts()],
});

require('esbuild').build({
  entryPoints: ['./src/index.js'],
  // outfile: "out.js",
  outdir: "dist/src",
  bundle: true,
  preserveSymlinks: true,
  minify: false,
  platform: "node",
}).catch(() => process.exit(1))


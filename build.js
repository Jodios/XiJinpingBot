require('esbuild').build({
  entryPoints: ['./src/index.js'],
  // outfile: "out.js",
  outdir: "dist/src",
  bundle: true,
  minify: true,
  platform: "node",
}).catch(() => process.exit(1))


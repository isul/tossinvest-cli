#!/usr/bin/env node
'use strict';

const { spawnSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const os = require('os');

const binaryName = process.platform === 'win32' ? 'tossinvest-cli.exe' : 'tossinvest-cli';

function findBinary() {
  const local = path.join(__dirname, binaryName);
  if (fs.existsSync(local)) {
    return local;
  }

  const goBin = path.join(process.env.GOPATH || path.join(os.homedir(), 'go'), 'bin', binaryName);
  if (fs.existsSync(goBin)) {
    return goBin;
  }

  const which = spawnSync(process.platform === 'win32' ? 'where' : 'which', [binaryName], {
    encoding: 'utf8',
  });
  if (which.status === 0) {
    return which.stdout.trim().split('\n')[0];
  }

  console.error(`tossinvest-cli binary not found.

Install via Go:
  go install github.com/isul/tossinvest-cli/cmd/tossinvest-cli@latest

Or build from source:
  git clone https://github.com/isul/tossinvest-cli.git
  cd tossinvest-cli && go build -o $(npm prefix -g)/lib/node_modules/@isul/tossinvest-cli/bin/${binaryName} ./cmd/tossinvest-cli/
`);
  process.exit(1);
}

const bin = findBinary();
const result = spawnSync(bin, process.argv.slice(2), { stdio: 'inherit' });
process.exit(result.status ?? 1);

import {
  WhirlpoolContext,
  ORCA_WHIRLPOOL_PROGRAM_ID,
} from '@orca-so/whirlpools-sdk';
import { Wallet, AnchorProvider } from '@coral-xyz/anchor';
import { Connection, PublicKey } from '@solana/web3.js';
import { loadKeypairFromFile } from '../utils/load';
import { getLiquidity } from '../func/getPool';
import { getSwapQuote } from '../func/getSwapQuote';

async function main() {
  // 1. Set up RPC connection
  const RPC_ENDPOINT = 'https://api.devnet.solana.com';
  const connection = new Connection(RPC_ENDPOINT, 'confirmed');

  // 2. Load wallet from local secret key file
  const SECRET_KEY_PATH = 'id.json';
  const wallet = new Wallet(loadKeypairFromFile(SECRET_KEY_PATH));

  // 3. Create Anchor provider
  const provider = new AnchorProvider(connection, wallet, {});

  // 4. Create Whirlpool context using the Devnet program ID
  const ctx = WhirlpoolContext.withProvider(
    provider,
    ORCA_WHIRLPOOL_PROGRAM_ID,
  );

  const poolAddress = new PublicKey(
    'CmH9ZMZCg2E3srXF4KR2j66A335wguNGYe1A5NFdpYUq',
  );

  await getLiquidity(ctx, poolAddress);
  await getSwapQuote(ctx, poolAddress);
}
void main();

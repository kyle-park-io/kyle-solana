import {
  buildWhirlpoolClient,
  WhirlpoolContext,
} from '@orca-so/whirlpools-sdk';
import { getAccount } from '@solana/spl-token';
import { PublicKey } from '@solana/web3.js';

/**
 * Fetch liquidity-related info from an Orca Whirlpool pool
 *
 * @param ctx - Whirlpool context with connection and wallet
 * @param poolAddress - The Whirlpool pool's public address
 */
export async function getLiquidity(
  ctx: WhirlpoolContext,
  poolAddress: PublicKey,
) {
  // Initialize the Whirlpool client
  const client = buildWhirlpoolClient(ctx);

  const whirlpool = await client.getPool(poolAddress);
  if (!whirlpool) {
    throw new Error('Failed to fetch Whirlpool pool information.');
  }

  // Retrieve token and vault information
  const tokenA = whirlpool.getTokenAInfo();
  const tokenB = whirlpool.getTokenBInfo();
  const vaultA = whirlpool.getTokenVaultAInfo();
  const vaultB = whirlpool.getTokenVaultBInfo();

  const conn = ctx.connection;

  // Get actual vault token balances using SPL Token getAccount
  const vaultAAccount = await getAccount(conn, vaultA.address);
  const vaultBAccount = await getAccount(conn, vaultB.address);

  const amountA = Number(vaultAAccount.amount) / 10 ** tokenA.decimals;
  const amountB = Number(vaultBAccount.amount) / 10 ** tokenB.decimals;

  console.log(`Token A (${tokenA.mint.toBase58()}): ${amountA}`);
  console.log(`Token B (${tokenB.mint.toBase58()}): ${amountB}`);

  return {
    tokenA: {
      mint: tokenA.mint.toBase58(),
      amount: amountA,
    },
    tokenB: {
      mint: tokenB.mint.toBase58(),
      amount: amountB,
    },
  };
}

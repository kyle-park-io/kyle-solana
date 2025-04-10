import {
  buildWhirlpoolClient,
  WhirlpoolContext,
  swapQuoteByInputToken,
  ORCA_WHIRLPOOL_PROGRAM_ID,
} from '@orca-so/whirlpools-sdk';
import { Percentage } from '@orca-so/common-sdk';
import { BN } from '@coral-xyz/anchor';
import { PublicKey } from '@solana/web3.js';

/**
 * Fetch a swap quote from the Orca Whirlpool SDK
 *
 * @param ctx - Whirlpool context containing connection and wallet
 * @param poolAddress - Address of the target Whirlpool pool
 */
export async function getSwapQuote(
  ctx: WhirlpoolContext,
  poolAddress: PublicKey,
) {
  // Initialize the Whirlpool client
  const client = buildWhirlpoolClient(ctx);

  const whirlpool = await client.getPool(poolAddress);
  if (!whirlpool) {
    throw new Error('Failed to fetch Whirlpool pool information.');
  }

  // Define input token and swap parameters
  const inputTokenMint = whirlpool.getTokenAInfo().mint;
  const inputAmount = new BN(100); // Example amount (in smallest unit, e.g., lamports)
  const slippage = Percentage.fromFraction(new BN(1), new BN(16)); // 6.25% slippage tolerance

  console.log('slippage:', slippage);

  // Get a quote for swapping inputToken with specified amount and slippage
  const quote = await swapQuoteByInputToken(
    whirlpool,
    inputTokenMint,
    inputAmount,
    slippage,
    ORCA_WHIRLPOOL_PROGRAM_ID,
    client.getFetcher(),
  );

  // Log the results
  console.log('Estimated output:', quote.estimatedAmountOut.toString());
  console.log(
    'Minimum output after slippage:',
    quote.otherAmountThreshold.toString(),
  );
  console.log(
    'Estimated sqrtPrice after swap:',
    quote.estimatedEndSqrtPrice.toString(),
  );
}

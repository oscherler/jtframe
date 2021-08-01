/*  This file is part of JT_FRAME.
    JTFRAME program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    JTFRAME program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with JTFRAME.  If not, see <http://www.gnu.org/licenses/>.

    Author: Jose Tejada Gomez. Twitter: @topapate
    Version: 1.0
    Date: 1-8-2021 */

module jtframe_hsize #( parameter
    WCOL        =4,    // bits per colour
    SCREEN_WIDTH=384   // screen pixel width (including blanking)
) (
    input             clk,
    input             pxl_cen,
    input             pxl2_cen,

    input       [3:0] scale,
    input       [4:0] offset,
    input             enable,

    input  [WCOL-1:0] r_in,
    input  [WCOL-1:0] g_in,
    input  [WCOL-1:0] b_in,
    input             HS_in,
    input             VS_in,
    input             HB_in,
    input             VB_in,
    // filtered video
    output reg            HS_out,
    output reg            VS_out,
    output reg            HB_out,
    output reg            VB_out,
    output reg [WCOL-1:0] r_out,
    output reg [WCOL-1:0] g_out,
    output reg [WCOL-1:0] b_out

);

localparam VW = SCREEN_WIDTH <= 256 ? 8 : (SCREEN_WIDTH <= 512 ? 9 : 10);

wire [WCOL*3-1:0] rgb_out, rgb_in;
reg  [    VW-1:0] rdcnt, wrcnt;
reg  [       4:0] summand;
reg  [       5:0] sum;

reg        line=0, suml, over, passz, HSl;
wire       we;
wire [5:0] next_sum;

assign rgb_in   = {r_in, g_in, b_in};
assign we       = pxl_cen;
assign next_sum = sum + {1'b0, summand};

always @(posedge clk) if(pxl_cen) begin
    HSl <= HS_in;
    if( HS_in & ~HSl ) begin
        line  <= ~line;
        wrcnt <= 0;
    end else if(wrcnt!=SCREEN_WIDTH-1) begin
        wrcnt <= wrcnt + 1'd1;
    end

    {r_out,g_out,b_out} <= enable ? (over || !passz ? 0 : rgb_out) : rgb_in;
    HS_out <= HS_in;
    VS_out <= VS_in;
    HB_out <= HB_in; // not accurate when scaling is enabled
    VB_out <= VB_in;
end

always @(posedge clk) if(pxl2_cen) begin
    suml <= sum[5];
    if( HS_in & ~HSl ) begin
        rdcnt <= { {VW-5{offset[4]}}, offset };
        sum   <= 0;
        over  <= 0;
        passz <= 0; // passed zero, used to avoid setting "over" wrong
                    // when using negative offsets
    end else begin
        sum  <= next_sum;
        if( sum[5] != next_sum[5] && !over ) begin
            if( rdcnt==0 ) passz <= 1;
            if( rdcnt == SCREEN_WIDTH-1 && passz ) begin
                over <= 1;
            end else begin
                rdcnt <= rdcnt + 1'd1;
            end
        end
    end
end

always @(*) begin
    case( scale )
        0: summand = 5'b0_0001;
        1: summand = 5'b0_0011;
        2: summand = 5'b0_0101;
        3: summand = 5'b0_0111;
        4: summand = 5'b0_1001;
        5: summand = 5'b0_1011;
        6: summand = 5'b0_1101;
        7: summand = 5'b0_1111;
        8: summand = 5'b1_0000; // 1:1
        9: summand = 5'b1_0001;
       10: summand = 5'b1_0010;
       11: summand = 5'b1_0011;
       12: summand = 5'b1_0100;
       13: summand = 5'b1_0101;
       14: summand = 5'b1_0110;
       15: summand = 5'b1_0111;
    endcase // scale
end

jtframe_dual_ram #(.dw(WCOL*3), .aw(VW+1)) u_line(
    .clk0   ( clk       ),
    .clk1   ( clk       ),
    // Port 0: writes
    .data0  ( rgb_in    ),
    .addr0  ({line, wrcnt}),
    .we0    ( pxl_cen   ),
    .q0     (           ),
    // Port 1
    .data1  (           ),
    .addr1  ({line,rdcnt}),
    .we1    ( 1'b0      ),
    .q1     ( rgb_out   )
);

endmodule
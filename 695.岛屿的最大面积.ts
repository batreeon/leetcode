/*
 * @lc app=leetcode.cn id=695 lang=typescript
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
// 并查集
function maxAreaOfIsland(grid: number[][]): number {
    let nr = grid.length;
    if (nr <= 0) {
        return 0;
    }
    let nc = grid[0].length;
    if (nc <= 0) {
        return 0;
    }

    let uf = new unionFind695(grid);
    for (let i = 0; i < nr; i++) {
        for (let j = 0; j < nc; j++) {
            if (grid[i][j] === 1 && i-1 >= 0 && grid[i-1][j] === 1) {
                uf.union((i-1)*nc+j, i*nc+j);
            }
            if (grid[i][j] === 1 && j-1 >= 0 && grid[i][j-1] === 1) {
                uf.union(i*nc+j-1, i*nc+j);
            }
        }
    }
    if (uf.rank.length === 0) {
        return 0;
    }
    //return uf.rank.sort((a,b) => (b - a))[0];
    return Math.max(...uf.rank);
};


class unionFind695 {
    private parent: number[] = new Array<number>();
    public rank: number[] = new Array<number>();

    constructor(grid: number[][]) {
        let nr = grid.length;
        let nc = grid[0].length;
        this.parent = new Array<number>(nr*nc).fill(-1);
        this.rank = new Array<number>(nr*nc).fill(0);
        for (let i = 0; i < nr; i++) {
            for (let j = 0; j < nc; j++) {
                if (grid[i][j] === 1) {
                    this.parent[i*nc+j] = i*nc+j;
                    this.rank[i*nc+j] = 1;
                }
            }
        }
    }

    private find(i: number): number {
        if (this.parent[i] === i) {
            return i;
        }
        return this.find(this.parent[i]);
    }

    public union(i: number, j: number) {
        let fi = this.find(i);
        let fj = this.find(j);
        if (fi !== fj) {
            if (this.rank[fi] < this.rank[fj]) {
                let temp: number = fi;
                fi = fj;
                fj = temp;
            }
            this.parent[fj] = fi;
            this.rank[fi] += this.rank[fj];
        }
    }
}
// @lc code=end


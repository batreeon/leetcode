/*
 * @lc app=leetcode.cn id=547 lang=typescript
 *
 * [547] 省份数量
 */

// @lc code=start
// 并查集，信手拈来
function findCircleNum(isConnected: number[][]): number {
    let uf = new unionFind(isConnected);
    for (let i = 0; i < isConnected.length; i++) {
        for (let j = 0; j < isConnected.length; j++) {
            if (isConnected[i][j] === 1) {
                uf.union(i,j);
            }
        }
    }
    return uf.getCount();
};

class unionFind {
    private parent: number[] = new Array<number>();
    private rank: number[] = new Array<number>();
    private count: number = 0;

    constructor(isConnected: number[][]) {
        let n: number = isConnected.length;
        for (let i = 0; i < n; i++) {
            this.parent[i] = i;
            this.rank[i] = 1;
            this.count++;
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
            if (this.rank[fi] === this.rank[fj]) {
                this.rank[fi]++;
            }
            this.count--;
        }
    }

    public getCount(): number {
        return this.count;
    }
}
// @lc code=end


/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Codec {
public:

    void rserialize(TreeNode* root,string& res) {
        if (!root) {
            res += "null,";
        }else{
            res += to_string(root->val)+",";
            rserialize(root->left,res);
            rserialize(root->right,res);
        }
    }
    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        string ret;
        rserialize(root,ret);
        return ret;
    }


    // Decodes your encoded data to tree.
    TreeNode* rdeserialize(list<string>& dataArray) {
        if (dataArray.front() == "null") {
            dataArray.erase(dataArray.begin());
            return nullptr;
        }
        TreeNode* root = new TreeNode(stoi(dataArray.front()));
        dataArray.erase(dataArray.begin());
        root->left = rdeserialize(dataArray);
        root->right = rdeserialize(dataArray);
        return root;
    }

    TreeNode* deserialize(string data) {
        list<string> dataArray;
        string str;
        for (auto &c : data) {
            if (c == ',') {
                dataArray.push_back(str);
                str.clear();
            }else{
                str.push_back(c);
            }
        } 
        return rdeserialize(dataArray);
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));
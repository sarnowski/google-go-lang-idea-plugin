package ro.redeul.google.go.lang.psi.impl.declarations;

import com.intellij.lang.ASTNode;
import com.intellij.psi.PsiElement;
import com.intellij.psi.ResolveState;
import com.intellij.psi.scope.PsiScopeProcessor;
import org.jetbrains.annotations.NotNull;
import ro.redeul.google.go.lang.psi.declarations.GoConstDeclaration;
import ro.redeul.google.go.lang.psi.declarations.GoVarDeclaration;
import ro.redeul.google.go.lang.psi.expressions.GoExpr;
import ro.redeul.google.go.lang.psi.expressions.literals.GoIdentifier;
import ro.redeul.google.go.lang.psi.impl.GoPsiElementBase;

/**
 * Author: Toader Mihai Claudiu <mtoader@gmail.com>
 * <p/>
 * Date: 7/16/11
 * Time: 4:10 AM
 */
public class GoConstDeclarationImpl extends GoPsiElementBase implements GoConstDeclaration {

    public GoConstDeclarationImpl(@NotNull ASTNode node) {
        super(node);
    }

    @Override
    public GoIdentifier[] getIdentifiers() {
        return findChildrenByClass(GoIdentifier.class);
    }

    @Override
    @NotNull
    public GoExpr[] getExpressions() {
        return findChildrenByClass(GoExpr.class);
    }

    @Override
    public boolean processDeclarations(@NotNull PsiScopeProcessor processor, @NotNull ResolveState state, PsiElement lastParent, @NotNull PsiElement place) {
        return  processor.execute(this, state);
    }

}

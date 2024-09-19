export const CurrencyFormatter = ({ amount, currency }) => {
    const formatter = new Intl.NumberFormat('id-ID', {
        style: 'currency',
        currency: currency || 'IDR',
    });

    return <span>{formatter.format(amount)}</span>;
};
